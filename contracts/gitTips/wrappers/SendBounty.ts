import { Address, beginCell, Cell, Contract, contractAddress, ContractProvider, Sender, SendMode } from 'ton-core';

export type SendBountyConfig = {};

export function sendBountyConfigToCell(config: SendBountyConfig): Cell {
    return beginCell().endCell();
}

export class SendBounty implements Contract {
    constructor(readonly address: Address, readonly init?: { code: Cell; data: Cell }) {}

    static createFromAddress(address: Address) {
        return new SendBounty(address);
    }

    static createFromConfig(config: SendBountyConfig, code: Cell, workchain = 0) {
        const data = sendBountyConfigToCell(config);
        const init = { code, data };
        return new SendBounty(contractAddress(workchain, init), init);
    }

    async sendDeploy(provider: ContractProvider, via: Sender, value: bigint) {
        await provider.internal(via, {
            value,
            sendMode: SendMode.PAY_GAS_SEPARATLY,
            body: beginCell().endCell(),
        });
    }
}
