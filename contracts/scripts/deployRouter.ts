import { toNano } from 'ton-core';
import { Router } from '../wrappers/Router';
import { compile, NetworkProvider } from '@ton-community/blueprint';

export async function run(provider: NetworkProvider) {
    const router = provider.open(Router.createFromConfig({}, await compile('Router')));

    await router.sendDeploy(provider.sender(), toNano('0.05'));

    await provider.waitForDeploy(router.address);

    // run methods on `router`
}
