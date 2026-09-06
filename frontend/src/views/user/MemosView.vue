<template>
  <AppLayout>
    <TablePageLayout>
      <template #filters>
        <div class="flex flex-wrap items-center gap-3">
          <div class="flex-1 sm:max-w-64">
            <input
              v-model="searchQuery"
              type="text"
              :placeholder="t('memos.searchPlaceholder')"
              class="input"
              @input="handleSearch"
            />
          </div>
          <div class="flex flex-1 flex-wrap items-center justify-end gap-2">
            <button
              @click="loadMemos"
              :disabled="loading"
              class="btn btn-secondary"
              :title="t('common.refresh')"
            >
              <Icon name="refresh" size="md" :class="loading ? 'animate-spin' : ''" />
            </button>
            <button @click="openEditor()" class="btn btn-primary">
              <Icon name="plus" size="md" class="mr-1" />
              {{ t('memos.create') }}
            </button>
          </div>
        </div>
      </template>

      <template #table>
        <DataTable
          :columns="columns"
          :data="filteredMemos"
          :loading="loading"
          default-sort-key="updated_at"
          default-sort-order="desc"
        >
          <template #cell-title="{ value, row }">
            <div class="min-w-0">
              <div class="flex items-center gap-2">
                <svg v-if="row.pinned" class="h-4 w-4 shrink-0 text-amber-500" fill="currentColor" viewBox="0 0 20 20">
                  <path d="M9.828 1.172a4 4 0 00-5.656 5.656l4 4a1 1 0 001.414 0l4-4a1 1 0 000-1.414l-1.586-1.586 3.293-3.293a1 1 0 10-1.414-1.414L10.414 6.586 9 5.172l.828-.828a1 1 0 000-1.414z"/>
                </svg>
                <span class="truncate font-medium text-gray-900 dark:text-white">{{ value }}</span>
              </div>
              <div class="mt-1 flex items-center gap-2 text-xs text-gray-500 dark:text-dark-400">
                <span>#{{ row.id }}</span>
                <span class="text-gray-300 dark:text-dark-700">·</span>
                <span>{{ formatDateTime(row.created_at) }}</span>
              </div>
            </div>
          </template>

          <template #cell-content="{ value }">
            <p class="line-clamp-2 max-w-md text-sm text-gray-500 dark:text-dark-400">{{ stripMarkdown(value) }}</p>
          </template>

          <template #cell-pinned="{ value }">
            <span
              :class="[
                'badge',
                value ? 'badge-warning' : 'badge-gray'
              ]"
            >
              {{ value ? t('memos.pinned') : t('memos.unpinned') }}
            </span>
          </template>

          <template #cell-updated_at="{ value }">
            <span class="text-sm text-gray-500 dark:text-dark-400">{{ formatRelativeTime(value) }}</span>
          </template>

          <template #cell-actions="{ row }">
            <div class="flex items-center space-x-1">
              <button
                @click="openEditor(row)"
                class="flex flex-col items-center gap-0.5 rounded-lg p-1.5 text-gray-500 transition-colors hover:bg-gray-100 hover:text-gray-700 dark:hover:bg-dark-600 dark:hover:text-gray-300"
                :title="t('common.edit')"
              >
                <Icon name="edit" size="sm" />
              </button>
              <button
                @click="togglePin(row)"
                :title="row.pinned ? t('memos.unpin') : t('memos.pin')"
                class="flex flex-col items-center gap-0.5 rounded-lg p-1.5 text-gray-500 transition-colors hover:bg-amber-50 hover:text-amber-600 dark:hover:bg-amber-900/20 dark:hover:text-amber-400"
              >
                <svg class="h-4 w-4" :class="row.pinned ? 'text-amber-500' : ''" fill="currentColor" viewBox="0 0 20 20">
                  <path d="M9.828 1.172a4 4 0 00-5.656 5.656l4 4a1 1 0 001.414 0l4-4a1 1 0 000-1.414l-1.586-1.586 3.293-3.293a1 1 0 10-1.414-1.414L10.414 6.586 9 5.172l.828-.828a1 1 0 000-1.414z"/>
                </svg>
              </button>
              <button
                @click="confirmDelete(row)"
                class="flex flex-col items-center gap-0.5 rounded-lg p-1.5 text-gray-500 transition-colors hover:bg-red-50 hover:text-red-600 dark:hover:bg-red-900/20 dark:hover:text-red-400"
                :title="t('common.delete')"
              >
                <Icon name="trash" size="sm" />
              </button>
            </div>
          </template>

          <template #empty>
            <EmptyState
              :title="t('empty.noData')"
              :description="t('memos.empty')"
              :action-text="t('memos.create')"
              @action="openEditor()"
            />
          </template>
        </DataTable>
      </template>
    </TablePageLayout>

    <!-- Editor Modal -->
    <Teleport to="body">
      <div v-if="editorOpen" class="fixed inset-0 z-[100] flex items-center justify-center p-4">
        <div class="absolute inset-0 bg-black/50" @click="closeEditor" />
        <div class="relative z-10 w-full max-w-2xl rounded-lg bg-white shadow-xl dark:bg-dark-800">
          <div class="flex items-center justify-between border-b border-gray-200 px-5 py-3 dark:border-dark-600">
            <h2 class="font-semibold text-gray-900 dark:text-white">
              {{ editingMemo ? t('memos.edit') : t('memos.create') }}
            </h2>
            <button @click="closeEditor" class="text-gray-400 hover:text-gray-600 dark:hover:text-dark-300">
              <Icon name="x" size="md" />
            </button>
          </div>
          <div class="space-y-4 p-5">
            <input
              v-model="form.title"
              type="text"
              :placeholder="t('memos.form.titlePlaceholder')"
              class="input"
              maxlength="200"
            />
            <textarea
              v-model="form.content"
              :placeholder="t('memos.form.contentPlaceholder')"
              rows="10"
              class="input resize-y font-mono text-sm"
            />
            <label class="flex items-center gap-2 text-sm text-gray-600 dark:text-dark-400">
              <input type="checkbox" v-model="form.pinned" class="checkbox" />
              {{ t('memos.form.pin') }}
            </label>
          </div>
          <div class="flex justify-end gap-2 border-t border-gray-200 px-5 py-3 dark:border-dark-600">
            <button @click="closeEditor" class="btn btn-secondary">{{ t('common.cancel') }}</button>
            <button @click="save" :disabled="saving" class="btn btn-primary">
              <Icon v-if="saving" name="refresh" size="sm" class="mr-1 animate-spin" />
              {{ t('common.save') }}
            </button>
          </div>
        </div>
      </div>
    </Teleport>

    <!-- Delete Confirmation -->
    <Teleport to="body">
      <div v-if="deleteTarget" class="fixed inset-0 z-[100] flex items-center justify-center p-4">
        <div class="absolute inset-0 bg-black/50" @click="deleteTarget = null" />
        <div class="relative z-10 w-full max-w-sm rounded-lg bg-white p-5 shadow-xl dark:bg-dark-800">
          <p class="text-gray-900 dark:text-white">{{ t('memos.deleteConfirm', { title: deleteTarget.title }) }}</p>
          <div class="mt-4 flex justify-end gap-2">
            <button @click="deleteTarget = null" class="btn btn-secondary">{{ t('common.cancel') }}</button>
            <button @click="doDelete" class="btn btn-danger">{{ t('common.delete') }}</button>
          </div>
        </div>
      </div>
    </Teleport>
  </AppLayout>
</template>

<script setup lang="ts">
import { ref, computed, onMounted } from 'vue'
import { useI18n } from 'vue-i18n'
import { useAppStore } from '@/stores/app'
import memosAPI from '@/api/memos'
import type { Memo } from '@/api/memos'
import type { Column } from '@/components/common/types'
import { formatDateTime } from '@/utils/format'

import AppLayout from '@/components/layout/AppLayout.vue'
import TablePageLayout from '@/components/layout/TablePageLayout.vue'
import DataTable from '@/components/common/DataTable.vue'
import EmptyState from '@/components/common/EmptyState.vue'
import Icon from '@/components/icons/Icon.vue'

const { t, locale } = useI18n()
const appStore = useAppStore()

const loading = ref(true)
const saving = ref(false)
const memos = ref<Memo[]>([])
const editorOpen = ref(false)
const editingMemo = ref<Memo | null>(null)
const deleteTarget = ref<Memo | null>(null)
const searchQuery = ref('')

const form = ref({
  title: '',
  content: '',
  pinned: false
})

const columns = computed<Column[]>(() => [
  { key: 'title', label: t('memos.columns.title'), sortable: true },
  { key: 'content', label: t('memos.columns.content') },
  { key: 'pinned', label: t('memos.columns.pinned'), sortable: true },
  { key: 'updated_at', label: t('memos.columns.updatedAt'), sortable: true },
  { key: 'actions', label: t('common.actions') }
])

const filteredMemos = computed(() => {
  const q = searchQuery.value.trim().toLowerCase()
  if (!q) return memos.value
  return memos.value.filter(
    (m) => m.title.toLowerCase().includes(q) || m.content.toLowerCase().includes(q)
  )
})

async function loadMemos() {
  loading.value = true
  try {
    memos.value = await memosAPI.list(100, 0)
  } catch {
    appStore.showError(t('memos.loadFailed'))
  } finally {
    loading.value = false
  }
}

let searchDebounceTimer: number | null = null
function handleSearch() {
  if (searchDebounceTimer) window.clearTimeout(searchDebounceTimer)
  searchDebounceTimer = window.setTimeout(() => {
    // client-side filter via computed, no action needed
  }, 300)
}

function openEditor(memo?: Memo) {
  editingMemo.value = memo ?? null
  form.value = {
    title: memo?.title ?? '',
    content: memo?.content ?? '',
    pinned: memo?.pinned ?? false
  }
  editorOpen.value = true
}

function closeEditor() {
  editorOpen.value = false
  editingMemo.value = null
}

async function save() {
  if (!form.value.title.trim() || !form.value.content.trim()) {
    appStore.showError(t('memos.form.required'))
    return
  }
  saving.value = true
  try {
    if (editingMemo.value) {
      await memosAPI.update(editingMemo.value.id, form.value)
      appStore.showSuccess(t('memos.updated'))
    } else {
      await memosAPI.create(form.value)
      appStore.showSuccess(t('memos.created'))
    }
    closeEditor()
    await loadMemos()
  } catch {
    appStore.showError(t('memos.saveFailed'))
  } finally {
    saving.value = false
  }
}

async function togglePin(memo: Memo) {
  try {
    await memosAPI.update(memo.id, { pinned: !memo.pinned })
    await loadMemos()
  } catch {
    appStore.showError(t('memos.saveFailed'))
  }
}

function confirmDelete(memo: Memo) {
  deleteTarget.value = memo
}

async function doDelete() {
  if (!deleteTarget.value) return
  try {
    await memosAPI.remove(deleteTarget.value.id)
    appStore.showSuccess(t('memos.deleted'))
    deleteTarget.value = null
    await loadMemos()
  } catch {
    appStore.showError(t('memos.deleteFailed'))
  }
}

function stripMarkdown(text: string): string {
  return text.replace(/[#*`>\-_~]/g, '').replace(/\n+/g, ' ').trim()
}

function formatRelativeTime(iso: string): string {
  const date = new Date(iso)
  const now = new Date()
  const diff = now.getTime() - date.getTime()
  const minutes = Math.floor(diff / 60000)
  const hours = Math.floor(diff / 3600000)
  const days = Math.floor(diff / 86400000)
  if (locale.value === 'zh') {
    if (days > 0) return `${days} 天前`
    if (hours > 0) return `${hours} 小时前`
    if (minutes > 0) return `${minutes} 分钟前`
    return '刚刚'
  }
  if (days > 0) return `${days}d ago`
  if (hours > 0) return `${hours}h ago`
  if (minutes > 0) return `${minutes}m ago`
  return 'just now'
}

onMounted(loadMemos)
</script>
