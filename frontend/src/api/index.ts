import type { AxiosProgressEvent, GenericAbortSignal } from 'axios'
// import { GetApi } from '../../wailsjs/go/main/App'
import { post } from '@/utils/request'
import { useSettingStore } from '@/store'

const endpoint = 'http://127.0.0.1:8080'

export function fetchChatAPI<T = any>(
  prompt: string,
  options?: { conversationId?: string; parentMessageId?: string },
  signal?: GenericAbortSignal,
) {
  return post<T>({
    url: `${endpoint}/chat`,
    data: { prompt, options },
    signal,
  })
}

export function fetchChatConfig<T = any>() {
  return post<T>({
    url: `${endpoint}/config`,
  })
}

export function fetchChatAPIProcess<T = any>(
  params: {
    prompt: string
    options?: { conversationId?: string; parentMessageId?: string }
    signal?: GenericAbortSignal
    onDownloadProgress?: (progressEvent: AxiosProgressEvent) => void },
) {
  const settingStore = useSettingStore()

  return post<T>({
    url: `${endpoint}/chat-process`,
    data: { prompt: params.prompt, options: params.options, systemMessage: settingStore.systemMessage },
    signal: params.signal,
    onDownloadProgress: params.onDownloadProgress,
  })
}

export function fetchSession<T>() {
  return post<T>({
    url: `${endpoint}/session`,
  })
}

export function fetchVerify<T>(token: string) {
  return post<T>({
    url: `${endpoint}/verify`,
    data: { token },
  })
}
