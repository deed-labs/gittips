import { Blockchain, SandboxContract } from '@ton-community/sandbox';
import { Cell, toNano } from 'ton-core';
import { Router } from '../wrappers/Router';
import '@ton-community/test-utils';
import { compile } from '@ton-community/blueprint';

describe('Router', () => {
    let code: Cell;

    beforeAll(async () => {
        code = await compile('Router');
    });

    let blockchain: Blockchain;
    let router: SandboxContract<Router>;

    beforeEach(async () => {
        blockchain = await Blockchain.create();

        router = blockchain.openContract(Router.createFromConfig({}, code));

        const deployer = await blockchain.treasury('deployer');

        const deployResult = await router.sendDeploy(deployer.getSender(), toNano('0.05'));

        expect(deployResult.transactions).toHaveTransaction({
            from: deployer.address,
            to: router.address,
            deploy: true,
        });
    });

    it('should deploy', async () => {
        // the check is done inside beforeEach
        // blockchain and router are ready to use
    });
});

