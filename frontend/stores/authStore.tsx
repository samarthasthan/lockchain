import { create } from "zustand";
import Web3 from "web3";

type AuthStore = {
  account: string;
  isConnected: boolean;
  web3: Web3 | null;
  setAccount: (account: string) => void;
  connectToMetaMask: () => Promise<void>;
  disconnect: () => void;
};

const authStore = create<AuthStore>((set) => ({
  account: "",
  isConnected: false,
  web3: null,

  setAccount: (account: string) =>
    set(() => ({
      account,
      isConnected: !!account,
    })),

  connectToMetaMask: async () => {
    const { ethereum } = window as any;

    if (typeof ethereum !== "undefined") {
      try {
        const web3 = new Web3(ethereum);
        const accounts = await web3.eth.getAccounts();
        if (accounts.length > 0) {
          set(() => ({
            account: accounts[0],
            isConnected: true,
            web3,
          }));
        } else {
          console.error("No accounts found.");
        }
      } catch (error) {
        console.error("Error connecting to MetaMask:", error);
      }
    } else {
      alert("Please install MetaMask!");
    }
  },

  disconnect: () =>
    set(() => ({
      account: "",
      isConnected: false,
      web3: null,
    })),
}));

export default authStore;
