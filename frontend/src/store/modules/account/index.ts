import { defineStore } from 'pinia'
import type { AccountInfo, AccountState, AccountType } from './helper'
import { defaultAccount, getLocalState, setLocalState } from './helper'

export const useAccountStore = defineStore('account-store', {
  state: (): AccountState => getLocalState(),
  actions: {
    updateAccountInfo(accountInfo: Partial<AccountInfo>) {
      this.accountInfo = { ...this.accountInfo, ...accountInfo }
      this.recordState()
    },

    resetAccountInfo() {
      this.accountInfo = { ...defaultAccount().accountInfo }
      this.recordState()
    },

    setAccountType(accountType: AccountType) {
      if (this.accountInfo.currentType !== accountType) {
        this.accountInfo.currentType = accountType
        this.recordState()
      }
    },

    recordState() {
      setLocalState(this.$state)
    },
  },
})
