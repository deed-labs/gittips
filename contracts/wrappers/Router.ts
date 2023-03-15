import {
    Address,
    beginCell,
    Cell,
    Contract,
    contractAddress,
    ContractProvider,
    Sender,
    SendMode,
    toNano,
} from 'ton-core';
import fs from 'fs';

const routerJSON = JSON.parse(fs.readFileSync('./build/Router.compiled.json').toString());

export class Router implements Contract {
    static readonly code = Cell.fromBoc(Buffer.from(routerJSON.hex, 'hex'))[0];

    readonly address: Address;
    readonly init: { code: Cell; data: Cell };

    constructor(
        workchain: number,
        initParams: {
            feeRate: number;
            adminAddr: Address;
            feeAddr: Address;
            budgetCode: Cell;
        }
    ) {
        const data = beginCell()
            .storeUint(initParams.feeRate, 8)
            .storeAddress(initParams.adminAddr)
            .storeAddress(initParams.feeAddr)
            .storeRef(initParams.budgetCode)
            .endCell();

        this.init = { code: Router.code, data };
        this.address = contractAddress(workchain, this.init);
    }

    async sendDeploy(provider: ContractProvider, via: Sender, value: bigint) {
        await provider.internal(via, {
            value,
            sendMode: SendMode.PAY_GAS_SEPARATELY,
            body: beginCell().endCell(),
        });
    }

    async sendAddBudget(
        provider: ContractProvider,
        via: Sender,
        params: {
            value: bigint;
        }
    ) {
        await provider.internal(via, {
            value: params.value,
            body: beginCell()
                .storeUint(1, 32) // op
                .storeUint(0, 64) // query id
                .endCell(),
        });
    }

    async sendWithdrawBudget(
        provider: ContractProvider,
        via: Sender,
        params: {
            value?: bigint;
            amount: bigint;
        }
    ) {
        await provider.internal(via, {
            value: params.value ?? toNano('0.05'),
            body: beginCell()
                .storeUint(2, 32) // op
                .storeUint(0, 64) // query id
                .storeCoins(params.amount) // amount
                .endCell(),
        });
    }

    async sendSendPayout(
        provider: ContractProvider,
        via: Sender,
        params: {
            value?: bigint;
            to_address: Address;
            amount: bigint;
        }
    ) {
        await provider.internal(via, {
            value: params.value ?? toNano('0.05'),
            body: beginCell()
                .storeUint(3, 32) // op
                .storeUint(0, 64) // query id
                .storeAddress(params.to_address) // to_address
                .storeCoins(params.amount) // amount
                .endCell(),
        });
    }

    async sendSetFee(
        provider: ContractProvider,
        via: Sender,
        params: {
            value?: bigint;
            feeRate: number;
        }
    ) {
        await provider.internal(via, {
            value: params.value ?? toNano('0.05'),
            body: beginCell()
                .storeUint(4, 32) // op
                .storeUint(0, 64) // query id
                .storeUint(params.feeRate, 8) // fee_rate
                .endCell(),
        });
    }

    async getBudgetAddress(provider: ContractProvider, address: Address) {
        const res = await provider.get('get_budget_address', [
            { type: 'slice', cell: beginCell().storeAddress(address).endCell() },
        ]);
        return res.stack.readAddress();
    }

    async getFeeRate(provider: ContractProvider) {
        const res = await provider.get('get_fee_rate', []);
        return res.stack.readNumber();
    }
}
