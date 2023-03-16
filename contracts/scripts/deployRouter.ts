import { Address, toNano } from 'ton-core';
import { Router } from '../wrappers/Router';
import { NetworkProvider } from '@ton-community/blueprint';
import { Budget } from '../wrappers/Budget';
import { load } from 'ts-dotenv';

const env = load({
    FEE_RATE: Number,
    ADMIN_ADDRESS: String,
    FEE_ADDRESS: String,
});

export async function run(provider: NetworkProvider) {
    let feeRate = env.FEE_RATE;
    if (!feeRate) throw new Error('fee rate not set');
    let adminAddress = env.ADMIN_ADDRESS;
    if (!adminAddress) throw new Error('admin address not set');
    let feeAddress = env.FEE_ADDRESS;
    if (!feeAddress) throw new Error('fee address not set');

    const router = provider.open(
        new Router(0, {
            feeRate: feeRate,
            adminAddr: Address.parse(adminAddress),
            feeAddr: Address.parse(feeAddress),
            budgetCode: Budget.code,
        })
    );

    await router.sendDeploy(provider.sender(), toNano('0.05'));

    await provider.waitForDeploy(router.address);
}
