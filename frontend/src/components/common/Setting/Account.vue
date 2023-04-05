<script lang="ts" setup>
import { computed, ref, watch } from 'vue'
import { NButton, NInput, NSelect, useMessage } from 'naive-ui'
import { SetAccountState } from '../../../../wailsjs/go/backend/Server'
import { backend } from '../../../../wailsjs/go/models'
import { useAccountStore } from '@/store'
import type { AccountInfo, AccountType } from '@/store/modules/account/helper'
import { t } from '@/locales'

const accountStore = useAccountStore()
const ms = useMessage()
const accountInfo = computed(() => accountStore.accountInfo)
const openaiApiKey = ref(accountInfo.value.openaiApiKey ?? '')
const openaiAccessToken = ref(accountInfo.value.openaiAccessToken ?? '')
const newbingCookies = ref(accountInfo.value.newbingCookies ?? '')
const proxy = ref(accountInfo.value.proxy ?? '')
// const baseURL = ref(accountInfo.value.baseURL ?? '')

const accountType = computed({
  get() {
    return accountStore.accountInfo.currentType
  },
  set(value: AccountType) {
    accountStore.setAccountType(value)
  },
})

const engineOptions: { label: string; key: AccountType; value: AccountType; disabled: boolean }[] = [
  { label: 'ChatGPT', key: 'ChatGPT', value: 'ChatGPT', disabled: false },
  { label: 'ChatGPTUnofficial', key: 'ChatGPTUnofficial', value: 'ChatGPTUnofficial', disabled: true },
  { label: 'NewBingUnofficial', key: 'NewBingUnofficial', value: 'NewBingUnofficial', disabled: false },
]

function updateAccountInfo(options: Partial<AccountInfo>) {
  accountStore.updateAccountInfo(options)
  ms.success(t('common.success'))
}

watch(accountInfo, (newValue, oldValue) => {
  if (newValue !== oldValue) {
    const accs = new backend.AccountState({
      account_info: new backend.AccountInfo({
        chat_engine: accountStore.accountInfo.currentType,
        openai_api_key: accountStore.accountInfo.openaiApiKey,
        base_url: accountStore.accountInfo.baseURL,
        openai_access_token: accountStore.accountInfo.openaiAccessToken,
        newbing_cookies: accountStore.accountInfo.newbingCookies,
        proxy: accountStore.accountInfo.proxy,
      }),
    })
    SetAccountState(accs).then(() => {
      console.log('watch accountStore: ', accountStore)
    })
  }
})
</script>

<template>
  <div class="p-4 space-y-5 min-h-[200px]">
    <div class="space-y-6">
      <!-- model -->
      <div class="flex items-center space-x-4">
        <span class="flex-shrink-0 w-[100px]">{{ $t('setting.engine') }}</span>
        <div class="flex flex-wrap items-center gap-4">
          <NSelect
            style="width: 240px"
            :value="accountType"
            :options="engineOptions"
            @update-value="value => updateAccountInfo({ currentType: value })"
          />
        </div>
      </div>

      <div v-if="accountType === 'ChatGPT'">
        <div class="flex items-center space-x-4">
          <span class="flex-shrink-0 w-[100px]">{{ $t('setting.openaiApiKey') }}</span>
          <div class="flex-1">
            <NInput v-model:value="openaiApiKey" placeholder="your openai api key" />
          </div>
          <NButton size="tiny" text type="primary" @click="updateAccountInfo({ openaiApiKey })">
            {{ $t('common.save') }}
          </NButton>
        </div>
      </div>

      <div v-if="accountType === 'ChatGPTUnofficial'">
        <div class="flex items-center space-x-4">
          <span class="flex-shrink-0 w-[100px]">{{ $t('setting.openaiToken') }}</span>
          <div class="flex-1">
            <NInput v-model:value="openaiAccessToken" placeholder="your openai access token" />
          </div>
          <NButton size="tiny" text type="primary" @click="updateAccountInfo({ openaiAccessToken })">
            {{ $t('common.save') }}
          </NButton>
        </div>
      </div>

      <div v-if="accountType === 'NewBingUnofficial'">
        <div class="flex items-center space-x-4">
          <span class="flex-shrink-0 w-[100px]">{{ $t('setting.newBingCookies') }}</span>
          <div class="flex-1">
            <NInput v-model:value="newbingCookies" placeholder="your new bing cookies" />
          </div>
          <NButton size="tiny" text type="primary" @click="updateAccountInfo({ newbingCookies })">
            {{ $t('common.save') }}
          </NButton>
        </div>
      </div>
      <!-- proxy -->
      <div class="flex items-center space-x-4">
        <span class="flex-shrink-0 w-[100px]">{{ $t('setting.proxy') }}</span>
        <div class="flex-1">
          <NInput v-model:value="proxy" placeholder="http://127.0.0.1:10809 or socks5://" />
        </div>
        <NButton size="tiny" text type="primary" @click="updateAccountInfo({ proxy })">
          {{ $t('common.save') }}
        </NButton>
      </div>
    </div>
  </div>
</template>
