<template>
  <!-- Row 1: Core Stats -->
  <div class="grid grid-cols-2 gap-4">
    <!-- API Keys -->
    <div class="card p-4">
      <div class="flex items-center gap-3">
        <div class="rounded-lg bg-blue-100 p-2 dark:bg-blue-900/30">
          <Icon name="key" size="md" class="text-blue-600 dark:text-blue-400" :stroke-width="2" />
        </div>
        <div>
          <p class="text-xs font-medium text-gray-500 dark:text-gray-400">{{ t('dashboard.apiKeys') }}</p>
          <p class="text-xl font-bold text-gray-900 dark:text-white">{{ stats?.total_api_keys || 0 }}</p>
          <p class="text-xs text-green-600 dark:text-green-400">{{ stats?.active_api_keys || 0 }} {{ t('common.active') }}</p>
        </div>
      </div>
    </div>

    <!-- Today Requests -->
    <div class="card p-4">
      <div class="flex items-center gap-3">
        <div class="rounded-lg bg-green-100 p-2 dark:bg-green-900/30">
          <Icon name="chart" size="md" class="text-green-600 dark:text-green-400" :stroke-width="2" />
        </div>
        <div>
          <p class="text-xs font-medium text-gray-500 dark:text-gray-400">{{ t('dashboard.todayRequests') }}</p>
          <p class="text-xl font-bold text-gray-900 dark:text-white">{{ stats?.today_requests || 0 }}</p>
          <p class="text-xs text-gray-500 dark:text-gray-400">{{ t('common.total') }}: {{ formatNumber(stats?.total_requests || 0) }}</p>
        </div>
      </div>
    </div>
  </div>

  <!-- Row 2: Token & Performance Stats -->
  <div class="grid grid-cols-2 gap-4 lg:grid-cols-4">
    <!-- Today Tokens -->
    <div class="card p-4">
      <div class="flex items-center gap-3">
        <div class="rounded-lg bg-amber-100 p-2 dark:bg-amber-900/30">
          <Icon name="cube" size="md" class="text-amber-600 dark:text-amber-400" :stroke-width="2" />
        </div>
        <div>
          <p class="text-xs font-medium text-gray-500 dark:text-gray-400">{{ t('dashboard.todayTokens') }}</p>
          <p class="text-xl font-bold text-gray-900 dark:text-white">{{ formatTokens(stats?.today_tokens || 0) }}</p>
        </div>
      </div>
    </div>

    <!-- Total Tokens -->
    <div class="card p-4">
      <div class="flex items-center gap-3">
        <div class="rounded-lg bg-indigo-100 p-2 dark:bg-indigo-900/30">
          <Icon name="database" size="md" class="text-indigo-600 dark:text-indigo-400" :stroke-width="2" />
        </div>
        <div>
          <p class="text-xs font-medium text-gray-500 dark:text-gray-400">{{ t('dashboard.totalTokens') }}</p>
          <p class="text-xl font-bold text-gray-900 dark:text-white">{{ formatTokens(stats?.total_tokens || 0) }}</p>
        </div>
      </div>
    </div>

    <!-- Performance (RPM/TPM) -->
    <div class="card p-4">
      <div class="flex items-center gap-3">
        <div class="rounded-lg bg-violet-100 p-2 dark:bg-violet-900/30">
          <Icon name="bolt" size="md" class="text-violet-600 dark:text-violet-400" :stroke-width="2" />
        </div>
        <div class="flex-1">
          <p class="text-xs font-medium text-gray-500 dark:text-gray-400">{{ t('dashboard.performance') }}</p>
          <div class="flex items-baseline gap-2">
            <p class="text-xl font-bold text-gray-900 dark:text-white">{{ formatTokens(stats?.rpm || 0) }}</p>
            <span class="text-xs text-gray-500 dark:text-gray-400">RPM</span>
          </div>
          <div class="flex items-baseline gap-2">
            <p class="text-sm font-semibold text-violet-600 dark:text-violet-400">{{ formatTokens(stats?.tpm || 0) }}</p>
            <span class="text-xs text-gray-500 dark:text-gray-400">TPM</span>
          </div>
        </div>
      </div>
    </div>

    <!-- Avg Response Time -->
    <div class="card p-4">
      <div class="flex items-center gap-3">
        <div class="rounded-lg bg-rose-100 p-2 dark:bg-rose-900/30">
          <Icon name="clock" size="md" class="text-rose-600 dark:text-rose-400" :stroke-width="2" />
        </div>
        <div>
          <p class="text-xs font-medium text-gray-500 dark:text-gray-400">{{ t('dashboard.avgResponse') }}</p>
          <p class="text-xl font-bold text-gray-900 dark:text-white">{{ formatDuration(stats?.average_duration_ms || 0) }}</p>
          <p class="text-xs text-gray-500 dark:text-gray-400">{{ t('dashboard.averageTime') }}</p>
        </div>
      </div>
    </div>
  </div>

  <!-- Row 3: Per-platform breakdown -->
  <div v-if="!isSimple && platformCards.length > 0" class="card p-4">
    <div class="mb-3 flex items-center justify-between">
      <h3 class="text-sm font-semibold text-gray-900 dark:text-white">{{ t('dashboard.platformBreakdown') }}</h3>
      <span class="text-xs text-gray-500 dark:text-gray-400">
        {{ t('dashboard.platformCount', { count: platformCards.length }) }}
      </span>
    </div>
    <div class="grid grid-cols-1 gap-3 sm:grid-cols-2 lg:grid-cols-4">
      <div
        v-for="item in platformCards"
        :key="item.platform"
        :class="[
          'rounded-lg border p-3',
          'border-gray-200 dark:border-dark-600'
        ]"
      >
        <div class="flex items-center justify-between">
          <span class="text-sm font-semibold text-gray-900 dark:text-white">
            {{ platformLabel(item.platform) }}
          </span>
        </div>
        <div class="mt-2 space-y-1 text-xs">
          <div class="flex items-center justify-between">
            <span class="text-gray-500 dark:text-gray-400">{{ t('dashboard.requests') }}</span>
            <span class="font-mono text-gray-700 dark:text-gray-300">
              {{ item.total_requests > 0 ? formatNumber(item.total_requests) : '-' }}
            </span>
          </div>
          <div class="flex items-center justify-between">
            <span class="text-gray-500 dark:text-gray-400">{{ t('dashboard.tokens') }}</span>
            <span class="font-mono text-gray-700 dark:text-gray-300">
              {{ item.total_tokens > 0 ? formatTokens(item.total_tokens) : '-' }}
            </span>
          </div>
        </div>

      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { computed } from 'vue'
import { useI18n } from 'vue-i18n'
import Icon from '@/components/icons/Icon.vue'
import type { UserDashboardStats as UserStatsType } from '@/api/usage'

interface FusedPlatformCard {
  platform: string
  total_requests: number
  total_tokens: number
}

const props = defineProps<{
  stats: UserStatsType
  isSimple: boolean
}>()
const { t } = useI18n()

const PLATFORM_LABELS: Record<string, string> = {
  anthropic: 'Claude',
  openai: 'OpenAI',
  gemini: 'Gemini',
  antigravity: 'Antigravity'
}

const platformLabel = (p: string) => PLATFORM_LABELS[p] ?? p

const platformCards = computed<FusedPlatformCard[]>(() => {
  const byPlat = new Map<string, NonNullable<UserStatsType['by_platform']>[number]>()
  for (const item of props.stats?.by_platform ?? []) byPlat.set(item.platform, item)
  const platforms = new Set<string>(byPlat.keys())

  const PLATFORM_ORDER = ['anthropic', 'openai', 'gemini', 'antigravity', 'grok']
  const cards: FusedPlatformCard[] = []

  for (const p of platforms) {
    const stat = byPlat.get(p)
    cards.push({
      platform: p,
      total_requests: stat?.total_requests ?? 0,
      total_tokens: stat?.total_tokens ?? 0,
    })
  }

  // 排序：按 PLATFORM_ORDER，未知平台按名称排序
  cards.sort((a, b) => {
    const ai = PLATFORM_ORDER.indexOf(a.platform)
    const bi = PLATFORM_ORDER.indexOf(b.platform)
    if (ai === -1 && bi === -1) return a.platform.localeCompare(b.platform)
    if (ai === -1) return 1
    if (bi === -1) return -1
    return ai - bi
  })

  return cards
})

const formatNumber = (n: number) => n.toLocaleString()
const formatTokens = (t: number) => {
  if (t >= 1_000_000) return `${(t / 1_000_000).toFixed(1)}M`
  if (t >= 1000) return `${(t / 1000).toFixed(1)}K`
  return t.toString()
}
const formatDuration = (ms: number) => ms >= 1000 ? `${(ms / 1000).toFixed(2)}s` : `${ms.toFixed(0)}ms`
</script>
