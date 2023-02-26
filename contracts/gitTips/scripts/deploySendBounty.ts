import { toNano } from 'ton-core';
import { SendBounty } from '../wrappers/SendBounty';
import { compile, NetworkProvider } from '@ton-community/blueprint';

export async function run(provider: NetworkProvider) {
    const sendBounty = SendBounty.createFromConfig({}, await compile('SendBounty'));

    await provider.deploy(sendBounty, toNano('0.05'));

    const openedContract = provider.open(sendBounty);

    // run methods on `openedContract`
}
