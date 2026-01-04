<script lang="ts">
  import { onMount } from 'svelte'
  import FormButton from '../form/Button.svelte'
  import FormInput from '../form/Input.svelte'
  import FormTextarea from '../form/Textarea.svelte'
  import { loadData, saveData } from '$lib/utils/apiHelpers'
  import { translate } from '$lib/i18n'
  import type { LetterData, LetterContent } from '$lib/types/models'

  // Reactive translation function
  let t = $derived($translate)

  interface Props {
    name: string
    legend: Record<string, string>
    onsend?: (name: string) => void
    onclose?: () => void
  }

  let { name, legend, onsend, onclose }: Props = $props()

  interface SettingResponse {
    id?: string
    key?: string
    value?: string | LetterContent
    [key: string]: unknown
  }

  interface LetterState extends LetterContent {
    id: string
    key: string
  }

  let letter = $state<LetterState>({
    id: '',
    key: '',
    subject: '',
    text: '',
    html: ''
  })
  let loading = $state(true)

  onMount(async () => {
    await loadLetter()
  })

  async function loadLetter() {
    if (!name) return

    loading = true
    const result = await loadData<SettingResponse | Record<string, SettingResponse>>(
      `/api/_/settings/${name}`,
      t('letter.failedToLoadLetter')
    )

    if (result) {
      const setting =
        (result as Record<string, SettingResponse>)[name] ||
        ((result as SettingResponse).id
          ? (result as SettingResponse)
          : Object.values(result as Record<string, SettingResponse>)[0])

      if (setting) {
        letter.id = setting.id || ''
        letter.key = setting.key || name

        if (setting.value) {
          const value = typeof setting.value === 'string' ? JSON.parse(setting.value) : setting.value
          letter.subject = value.subject || ''
          letter.text = value.text || ''
          letter.html = value.html || ''
        }
      }
    }
    loading = false
  }

  async function updateLetter() {
    const value: LetterContent = {
      subject: letter.subject,
      text: letter.text,
      html: letter.html
    }

    const update: LetterData = {
      id: letter.id,
      key: letter.key,
      value: JSON.stringify(value)
    }

    await saveData<LetterData>(`/api/_/settings/${name}`, update, true, t('letter.letterUpdated'), t('letter.failedToUpdateLetter'))
  }

  function handleSend() {
    onsend?.(name)
  }

  function close() {
    onclose?.()
  }

  function getTemplateKey(key: string): string {
    return `{{.${key}}}`
  }
</script>

<div>
  <div class="pb-8">
    <div class="flex items-center">
      <div class="pr-3">
        <h1>{t('letter.updateLetter')}</h1>
      </div>
    </div>
  </div>

  {#if loading}
    <div class="py-8 text-center">{t('common.loading')}</div>
  {:else}
    <div class="flow-root">
      <div class="flow-root">
        <dl class="mx-auto -my-3 mt-2 mb-0 space-y-4 text-sm">
          <FormInput id="subject" type="text" title={t('letter.subject')} bind:value={letter.subject} onfocusout={updateLetter} />
        </dl>
      </div>

      <dl class="mx-auto -my-3 mt-5 mb-0 space-y-4 text-sm">
        <FormTextarea id="textarea" title={t('letter.message')} bind:value={letter.text} rows={15} onfocusout={updateLetter} />
      </dl>
    </div>
  {/if}

  <div class="pt-8">
    <div class="flex">
      <div class="flex-none">
        <FormButton type="button" name={t('common.close')} color="gray" onclick={close} />
      </div>
      <div class="grow"></div>
      <div class="flex-none">
        <FormButton type="button" name={t('letter.testLetter')} color="cyan" onclick={handleSend} />
      </div>
    </div>
  </div>

  <table class="mt-8 text-base">
    <tbody>
      {#each Object.entries(legend) as [key, value] (key)}
        <tr class="cursor-default">
          <td class="w-32 font-bold">{getTemplateKey(key)}</td>
          <td>{value}</td>
        </tr>
      {/each}
    </tbody>
  </table>
</div>
