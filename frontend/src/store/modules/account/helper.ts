import { ss } from '@/utils/storage'

const LOCAL_NAME = 'accountStorage'

export type AccountType = 'ChatGPT' | 'ChatGPTUnofficial' | 'NewBingUnofficial'

export interface AccountInfo {
  currentType: AccountType
  openaiApiKey: string
  baseURL: string
  openaiAccessToken: string
  newbingCookies: string
  proxy: string
}

export interface AccountState {
  accountInfo: AccountInfo
}

export function defaultAccount(): AccountState {
  return {
    accountInfo: {
      currentType: 'ChatGPT',
      openaiApiKey: 'your openai api key',
      baseURL: '',
      openaiAccessToken: 'your openai access token',
      newbingCookies: 'your new bing cookies',
      proxy: '',
    },

  }
}

// setAccountType

export function getLocalState(): AccountState {
  const localAccount: AccountState | undefined = ss.get(LOCAL_NAME)
  return { ...defaultAccount(), ...localAccount }
}

export function setLocalState(account: AccountState): void {
  ss.set(LOCAL_NAME, account)
}

export function removeLocalState() {
  ss.remove(LOCAL_NAME)
}
