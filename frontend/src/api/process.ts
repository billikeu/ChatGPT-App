import type { AxiosProgressEvent, GenericAbortSignal } from 'axios'
import { DownloadProcess } from '../../wailsjs/go/main/App'

export function DownloadProcessDemo<T = any>(
  params: {
    prompt: string
    options?: { conversationId?: string; parentMessageId?: string }
    signal?: GenericAbortSignal
    onDownloadProgress?: (progressEvent: AxiosProgressEvent) => void },
) {
  // const settingStore = useSettingStore()
  return DownloadProcess('/chat-process', params.onDownloadProgress)
}
