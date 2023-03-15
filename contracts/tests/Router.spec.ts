import { Blockchain, SandboxContract, TreasuryContract } from '@ton-community/sandbox';
import { Cell, toNano } from 'ton-core';
import { Router } from '../wrappers/Router';
import '@ton-community/test-utils';
import { compile } from '@ton-community/blueprint';
import { Budget } from '../wrappers/Budget';

describe('Router', () => {
    let blockchain: Blockchain;
    let router: SandboxContract<Router>;

    let admin: SandboxContract<TreasuryContract>;

    beforeAll(async () => {
        blockchain = await Blockchain.create();

        admin = await blockchain.treasury('admin')

        router = blockchain.openContract(new Router(0,{
            feeRate: 10,
            adminAddr: admin.address,
            feeAddr: admin.address,
            budgetCode: Budget.code,
        }));

        const deployResult = await router.sendDeploy(admin.getSender(), toNano('0.05'));

        expect(deployResult.transactions).toHaveTransaction({
            from: admin.address,
            to: router.address,
            deploy: true,
        });
    });

    it('should add budget', async () => {

        const result = await router.addBudget(admin.getSender(), {
            value: toNano('0.05'),
        });

        expect(result.transactions).toHaveTransaction({
            from: admin.address,
            to: router.address,
            value: toNano('0.05'),
            body: expect.any(Cell),
        });
    });
});

