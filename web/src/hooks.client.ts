import { TON } from './lib/stores';
import { MyTonWallet, OpenMask, TonKeeper, TonSafe } from './lib/pkg/wallet/ton';

await TON.init({
	TonKeeper: new TonKeeper(),
	OpenMask: new OpenMask(),
	MyTonWallet: new MyTonWallet(),
	TonSafe: new TonSafe()
});
