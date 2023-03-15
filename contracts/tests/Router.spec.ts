import { Blockchain, SandboxContract, TreasuryContract } from '@ton-community/sandbox';
import { Cell, toNano } from 'ton-core';
import { Router } from '../wrappers/Router';
import '@ton-community/test-utils';
import { Budget } from '../wrappers/Budget';
import { inspect } from 'util';

describe('Router', () => {
    let blockchain: Blockchain;
    let router: SandboxContract<Router>;
    let admin: SandboxContract<TreasuryContract>;
    let owner: SandboxContract<TreasuryContract>;
    let random: SandboxContract<TreasuryContract>;

    beforeAll(async () => {
        blockchain = await Blockchain.create();

        admin = await blockchain.treasury('admin');
        owner = await blockchain.treasury('owner');
        random = await blockchain.treasury('random');

        router = blockchain.openContract(
            new Router(0, {
                feeRate: 10,
                adminAddr: admin.address,
                feeAddr: admin.address,
                budgetCode: Budget.code,
            })
        );

        const deployResult = await router.sendDeploy(admin.getSender(), toNano('0.05'));

        expect(deployResult.transactions).toHaveTransaction({
            from: admin.address,
            to: router.address,
            deploy: true,
        });
    });

    it('should return correct data', async () => {
        let feeRate = await router.getFeeRate();

        expect(feeRate).toEqual(10);
    });

    it('should set fee rate', async () => {
        await router.sendSetFee(admin.getSender(), {
            feeRate: 15,
        });
        let feeRate = await router.getFeeRate();

        expect(feeRate).toEqual(15);
    });

    it('should add budget', async () => {
        await router.sendAddBudget(owner.getSender(), {
            value: toNano('10.0'),
        });
        let budgetAddr = await router.getBudgetAddress(owner.address);

        expect(budgetAddr.toString()).toEqual(Budget.calculateAddress(router.address, owner.address).toString());

        let budget = await blockchain.getContract(budgetAddr);

        // Compare with 9.0 because of the fees
        // TODO: calculate fees and compare with the exact amount.
        expect(budget.balance).toBeGreaterThan(toNano('9.0'));
    });

    it('should withdraw budget', async () => {
        let ownerContract = await blockchain.getContract(owner.address);
        let oldOwnerBalance = ownerContract.balance;

        let result = await router.sendWithdrawBudget(owner.getSender(), {
            amount: toNano('5.0'),
        });

        let budgetAddr = await router.getBudgetAddress(owner.address);
        let budgetContract = await blockchain.getContract(budgetAddr);

        ownerContract = await blockchain.getContract(owner.address);
        let newOwnerBalance = ownerContract.balance;

        expect(result.transactions).toHaveTransaction({
            from: budgetAddr,
            to: owner.address,
            value: toNano('5.0'),
            success: true,
        });

        // Compare with ±0.1 TON to level paid fees
        // TODO: calculate fees and compare with the exact amount.
        expect(newOwnerBalance).toBeGreaterThan(toNano('4.9') + oldOwnerBalance);
        expect(budgetContract.balance).toBeLessThan(toNano('5.1'));
    });

    it('should send payout', async () => {
        let toContract = await blockchain.getContract(random.address);
        let oldToBalance = toContract.balance;

        let result = await router.sendSendPayout(admin.getSender(), {
            ownerAddress: owner.address,
            toAddress: toContract.address,
            amount: toNano('1.0'),
        });

        let budgetAddr = await router.getBudgetAddress(owner.address);
        let budgetContract = await blockchain.getContract(budgetAddr);

        toContract = await blockchain.getContract(toContract.address);
        let newToBalance = toContract.balance;

        console.log(inspect(result.transactions, false, 10000));

        expect(result.transactions).toHaveTransaction({
            from: budgetAddr,
            to: toContract.address,
            value: toNano('1.0'),
            success: true,
        });

        expect(result.transactions).toHaveTransaction({
            from: budgetAddr,
            to: admin.address,
            value: toNano('0.15'),
            success: true,
        });

        // Compare with ±0.1 TON to level paid fees
        // TODO: calculate fees and compare with the exact amount.
        expect(newToBalance).toBeGreaterThan(toNano('0.9') + oldToBalance);
        expect(budgetContract.balance).toBeLessThan(toNano('4.1'));
    });
});
