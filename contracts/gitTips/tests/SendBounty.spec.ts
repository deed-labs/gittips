import { Blockchain } from '@ton-community/sandbox';
import { Cell, toNano } from 'ton-core';
import { SendBounty } from '../wrappers/SendBounty';
import '@ton-community/test-utils';
import { compile } from '@ton-community/blueprint';

describe('SendBounty', () => {
    let code: Cell;

    beforeAll(async () => {
        code = await compile('SendBounty');
    });

    it('should deploy', async () => {
        const blockchain = await Blockchain.create();

        const sendBounty = blockchain.openContract(SendBounty.createFromConfig({}, code));

        const deployer = await blockchain.treasury('deployer');

        const deployResult = await sendBounty.sendDeploy(deployer.getSender(), toNano('0.05'));

        expect(deployResult.transactions).toHaveTransaction({
            from: deployer.address,
            to: sendBounty.address,
            deploy: true,
        });
    });
});
