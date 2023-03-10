/**
 * All wallet classes must implement this interface.
 */
export interface IWallet {
	/**
	 * True if user has wallet installed in browser, false otherwise.
	 * For external wallets its always true.
	 */
	available: boolean;
	/**
	 * The connection status.
	 * True if connected, false otherwise.
	 *
	 * @public
	 * @type {boolean}
	 */
	connected: boolean;
	/**
	 * The address of the wallet.
	 *
	 * @public
	 * @type {string}
	 */
	address: string;

	injected: boolean;
	embedded: boolean;

	/**
	 * Wallet constructor finalizer.
	 * It executes all asynchronous instructions necessary for the wallet to work properly.
	 */
	ready: Promise<void>;

	/**
	 * Returns link for connection through external wallet.
	 * Use it for generating QR code or some other way to connect.
	 * Callback function is called when connection is established. The address is passed as a parameter.
	 *
	 * @param {function} cb
	 * @returns {string} link
	 */
	connectExternal(cb: (address: string) => void): Promise<string>;

	connectInjected(cb: (address: string) => void): Promise<void>;

	disconnect(): Promise<void>;
}
