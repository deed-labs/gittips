import { TON } from './lib/stores/network';
import { MyTonWallet, OpenMask, TonKeeper, TonSafe } from './lib/pkg/wallet/ton';

// Wrap async calls with async function that executes itself to decrease the level of await.
// See https://github.com/vitejs/vite/issues/6985
(async () => {
	await TON.init({
		TonKeeper: new TonKeeper(),
		OpenMask: new OpenMask(),
		MyTonWallet: new MyTonWallet(),
		TonSafe: new TonSafe()
	});
})();
