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

    beforeAll(async () => {
        blockchain = await Blockchain.create();

        admin = await blockchain.treasury('admin');

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
        await router.sendAddBudget(admin.getSender(), {
            value: toNano('10.0'),
        });
        let budgetAddr = await router.getBudgetAddress(admin.address);

        expect(budgetAddr.toString()).toEqual(Budget.calculateAddress(router.address, admin.address).toString());

        let budget = await blockchain.getContract(budgetAddr);

        // Compare with 9.0 because of the fees
        // TODO: calculate fees and compare with the exact amount.
        expect(budget.balance).toBeGreaterThan(toNano('9.0'));
    });

    // FIXME
    // it('should withdraw budget', async () => {
    //     let res = await router.sendWithdrawBudget(admin.getSender(), {
    //         amount: toNano('5.0'),
    //     });

    //     console.log(inspect(res.transactions, false, 10000));

    //     console.log((await blockchain.getContract(router.address)).balance);

    //     let budgetAddr = await router.getBudgetAddress(admin.address);
    //     let budget = await blockchain.getContract(budgetAddr);

    //     // Compare with 5.0 because of the fees
    //     // TODO: calculate fees and compare with the exact amount.
    //     expect(budget.balance).toBeLessThan(toNano('5.0'));
    // });
});
