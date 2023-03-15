import { Blockchain, SandboxContract } from '@ton-community/sandbox';
import { Cell, toNano } from 'ton-core';
import { Budget } from '../wrappers/Budget';
import '@ton-community/test-utils';
import { compile } from '@ton-community/blueprint';

describe('Budget', () => {
    let code: Cell;

    beforeAll(async () => {
        code = await compile('Budget');
    });

    let blockchain: Blockchain;
    let budget: SandboxContract<Budget>;

    beforeEach(async () => {
        blockchain = await Blockchain.create();

        budget = blockchain.openContract(Budget.createFromConfig({}, code));

        const deployer = await blockchain.treasury('deployer');

        const deployResult = await budget.sendDeploy(deployer.getSender(), toNano('0.05'));

        expect(deployResult.transactions).toHaveTransaction({
            from: deployer.address,
            to: budget.address,
            deploy: true,
        });
    });

    it('should deploy', async () => {
        // the check is done inside beforeEach
        // blockchain and budget are ready to use
    });
});

