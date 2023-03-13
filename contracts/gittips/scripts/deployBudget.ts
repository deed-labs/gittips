import { toNano } from 'ton-core';
import { Budget } from '../wrappers/Budget';
import { compile, NetworkProvider } from '@ton-community/blueprint';

export async function run(provider: NetworkProvider) {
    const budget = provider.open(Budget.createFromConfig({}, await compile('Budget')));

    await budget.sendDeploy(provider.sender(), toNano('0.05'));

    await provider.waitForDeploy(budget.address);

    // run methods on `router`
}
