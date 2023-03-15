import { Address, beginCell, Cell, Contract, contractAddress, ContractProvider, Sender, SendMode } from 'ton-core';

export type BudgetConfig = {};

export function budgetConfigToCell(config: BudgetConfig): Cell {
    return beginCell().endCell();
}

export class Budget implements Contract {
    constructor(readonly address: Address, readonly init?: { code: Cell; data: Cell }) {}

    static createFromAddress(address: Address) {
        return new Budget(address);
    }

    static createFromConfig(config: BudgetConfig, code: Cell, workchain = 0) {
        const data = budgetConfigToCell(config);
        const init = { code, data };
        return new Budget(contractAddress(workchain, init), init);
    }

    async sendDeploy(provider: ContractProvider, via: Sender, value: bigint) {
        await provider.internal(via, {
            value,
            sendMode: SendMode.PAY_GAS_SEPARATELY,
            body: beginCell().endCell(),
        });
    }
}
