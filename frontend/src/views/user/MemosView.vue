<template>
  <AppLayout>
    <div class="mx-auto max-w-4xl px-4 py-6">
      <!-- Header -->
      <div class="mb-6 flex items-center justify-between">
        <h1 class="text-xl font-bold text-gray-900 dark:text-white">{{ t('memos.title') }}</h1>
        <button @click="openEditor()" class="btn btn-primary">
          <Icon name="plus" size="md" class="mr-1" />
          {{ t('memos.create') }}
        </button>
      </div>

      <!-- Loading -->
      <div v-if="loading" class="flex justify-center py-12">
        <Icon name="refresh" size="lg" class="animate-spin text-gray-400" />
      </div>

      <!-- Empty -->
      <div v-else-if="memos.length === 0" class="rounded-lg border border-dashed border-gray-300 py-16 text-center dark:border-dark-600">
        <p class="text-gray-400">{{ t('memos.empty') }}</p>
        <button @click="openEditor()" class="btn btn-primary mt-4">
          {{ t('memos.create') }}
        </button>
      </div>

      <!-- Memo List -->
      <div v-else class="space-y-3">
        <div
          v-for="memo in memos"
          :key="memo.id"
          class="group rounded-lg border border-gray-200 bg-white p-4 transition hover:shadow-md dark:border-dark-600 dark:bg-dark-800"
        >
          <div class="flex items-start justify-between gap-3">
            <div class="min-w-0 flex-1" @click="openEditor(memo)" class="cursor-pointer">
              <div class="flex items-center gap-2">
                <svg v-if="memo.pinned" class="h-4 w-4 shrink-0 text-amber-500" fill="currentColor" viewBox="0 0 20 20">
                  <path d="M9.828 1.172a4 4 0 00-5.656 5.656l4 4a1 1 0 001.414 0l4-4a1 1 0 000-1.414l-1.586-1.586 3.293-3.293a1 1 0 10-1.414-1.414L10.414 6.586 9 5.172l.828-.828a1 1 0 000-1.414z"/>
                </svg>
                <h3 class="truncate font-medium text-gray-900 dark:text-white">{{ memo.title }}</h3>
              </div>
              <p class="mt-1 line-clamp-2 text-sm text-gray-500 dark:text-dark-400">{{ stripMarkdown(memo.content) }}</p>
              <p class="mt-1 text-xs text-gray-400 dark:text-dark-500">{{ formatRelativeTime(memo.updated_at) }}</p>
            </div>
            <div class="flex shrink-0 items-center gap-1 opacity-0 transition group-hover:opacity-100">
              <button
                @click="togglePin(memo)"
                :title="memo.pinned ? t('memos.unpin') : t('memos.pin')"
                class="rounded p-1.5 text-gray-400 hover:bg-gray-100 hover:text-gray-600 dark:hover:bg-dark-700"
              >
                <svg class="h-4 w-4" :class="memo.pinned ? 'text-amber-500' : ''" fill="currentColor" viewBox="0 0 20 20">
                  <path d="M9.828 1.172a4 4 0 00-5.656 5.656l4 4a1 1 0 001.414 0l4-4a1 1 0 000-1.414l-1.586-1.586 3.293-3.293a1 1 0 10-1.414-1.414L10.414 6.586 9 5.172l.828-.828a1 1 0 000-1.414z"/>
                </svg>
              </button>
              <button
                @click="confirmDelete(memo)"
                :title="t('common.delete')"
                class="rounded p-1.5 text-gray-400 hover:bg-red-50 hover:text-red-600 dark:hover:bg-red-900/20"
              >
                <Icon name="trash" size="sm" />
              </button>
            </div>
          </div>
        </div>
      </div>
    </div>

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
import { ref, onMounted } from 'vue'
import { useI18n } from 'vue-i18n'
import { useAppStore } from '@/stores/app'
import memosAPI from '@/api/memos'
import type { Memo } from '@/api/memos'

const { t, locale } = useI18n()
const appStore = useAppStore()

const loading = ref(true)
const saving = ref(false)
const memos = ref<Memo[]>([])
const editorOpen = ref(false)
const editingMemo = ref<Memo | null>(null)
const deleteTarget = ref<Memo | null>(null)

const form = ref({
  title: '',
  content: '',
  pinned: false
})

async function loadMemos() {
  loading.value = true
  try {
    memos.value = await memosAPI.list()
  } catch {
    appStore.showError(t('memos.loadFailed'))
  } finally {
    loading.value = false
  }
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
