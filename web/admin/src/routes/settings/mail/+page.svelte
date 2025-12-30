<script lang="ts">
  import { onMount } from 'svelte';
  import Main from '$lib/layouts/Main.svelte';
  import Drawer from '$lib/components/Drawer.svelte';
  import Letter from '$lib/components/setting/Letter.svelte';
  import FormButton from '$lib/components/form/Button.svelte';
  import FormInput from '$lib/components/form/Input.svelte';
  import FormSelect from '$lib/components/form/Select.svelte';
  import { apiGet, apiUpdate, showMessage } from '$lib/utils';

  let smtp = {
    host: '',
    port: '',
    encryption: '',
    username: '',
    password: ''
  };
  let formErrors: Record<string, string> = {};
  let loading = true;
  let drawerOpen = false;
  let drawerMode: 'mail_letter_payment' | 'mail_letter_purchase' | null = null;

  const letterLegend = {
    mail_letter_payment: {
      Site_Name: 'Site name',
      Amount_Payment: 'Amount of payment',
      Payment_URL: 'Payment link'
    },
    mail_letter_purchase: {
      Purchases: 'Purchases',
      Admin_Email: 'Admin email'
    }
  };

  onMount(async () => {
    await loadSettings();
  });

  async function loadSettings() {
    loading = true;
    try {
      const res = await apiGet('/api/_/settings/mail');
      if (res.success) {
        smtp = {
          host: res.result.smtp?.host || '',
          port: res.result.smtp?.port || '',
          encryption: res.result.smtp?.encryption || '',
          username: res.result.smtp?.username || '',
          password: res.result.smtp?.password || ''
        };
      } else {
        showMessage(res.message || 'Failed to load settings', 'connextError');
      }
    } catch (error) {
      showMessage('Network error', 'connextError');
    } finally {
      loading = false;
    }
  }

  async function handleSubmit() {
    formErrors = {};

    if (!smtp.host || smtp.host.length < 4) {
      formErrors.smtp_host = 'SMTP host must be at least 4 characters';
    }
    if (!smtp.port || isNaN(Number(smtp.port))) {
      formErrors.smtp_port = 'SMTP port must be a number';
    }
    if (!smtp.encryption) {
      formErrors.smtp_encryption = 'Encryption is required';
    }
    if (!smtp.username) {
      formErrors.smtp_username = 'Username is required';
    }
    if (!smtp.password || smtp.password.length < 6) {
      formErrors.smtp_password = 'Password must be at least 6 characters';
    }

    if (Object.keys(formErrors).length > 0) {
      return;
    }

    try {
      const update = {
        smtp: {
          ...smtp,
          port: Number(smtp.port)
        }
      };
      const res = await apiUpdate('/api/_/settings/mail', update);
      if (res.success) {
        showMessage(res.message || 'Settings saved', 'connextSuccess');
      } else {
        showMessage(res.message || 'Failed to save settings', 'connextError');
      }
    } catch (error) {
      showMessage('Network error', 'connextError');
    }
  }

  async function sendTestLetter(letterName: string) {
    try {
      const res = await apiGet(`/api/_/test/letter/${letterName}`);
      if (res.success) {
        showMessage(res.message || 'Test letter sent', 'connextSuccess');
      } else {
        showMessage(res.message || 'Failed to send test letter', 'connextError');
      }
    } catch (error) {
      showMessage('Network error', 'connextError');
    }
  }

  function openDrawer(mode: 'mail_letter_payment' | 'mail_letter_purchase') {
    drawerMode = mode;
    drawerOpen = true;
  }

  function closeDrawer() {
    drawerOpen = false;
    setTimeout(() => {
      drawerMode = null;
    }, 200);
  }
</script>

<svelte:component this={Main}>
  <div class="pb-10">
    <header class="mb-4">
      <h1>Mail</h1>
    </header>

    <div>
      <h2 class="mb-5">Mail letters</h2>
      <div class="flex">
        <div 
          class="cursor-pointer rounded bg-gray-200 p-2" 
          on:click={() => openDrawer('mail_letter_payment')}
          role="button"
          tabindex="0"
          on:keydown={(e) => { if (e.key === 'Enter' || e.key === ' ') { e.preventDefault(); openDrawer('mail_letter_payment'); } }}
        >
          Letter of payment
        </div>
        <div 
          class="cursor-pointer rounded bg-gray-200 p-2 ml-5" 
          on:click={() => openDrawer('mail_letter_purchase')}
          role="button"
          tabindex="0"
          on:keydown={(e) => { if (e.key === 'Enter' || e.key === ' ') { e.preventDefault(); openDrawer('mail_letter_purchase'); } }}
        >
          Letter of purchase
        </div>
      </div>
      <hr class="mt-5" />
    </div>

    <div class="mt-5">
      <h2 class="mb-5">SMTP settings</h2>
      {#if !smtp.host || !smtp.port || !smtp.username || !smtp.password}
        <div class="mb-5 flex items-center justify-between bg-red-600 px-2 py-3 text-white">
          <p class="text-sm font-medium">This section is required!</p>
        </div>
      {/if}

      <form on:submit|preventDefault={handleSubmit}>
        <div class="flex">
          <div class="pr-3 w-64">
            <FormInput
              id="smtp_host"
              type="text"
              title="SMTP host"
              bind:value={smtp.host}
              error={formErrors.smtp_host}
              ico="server"
            />
          </div>
          <div class="pr-3 w-64">
            <FormInput
              id="smtp_port"
              type="text"
              title="SMTP port"
              bind:value={smtp.port}
              error={formErrors.smtp_port}
              ico="arrow-left-on-rectangle"
            />
          </div>
          <div>
            <FormSelect
              id="smtp_encryption"
              title="Encryption"
              options={['None', 'SSL/TLS', 'STARTTLS']}
              bind:value={smtp.encryption}
              error={formErrors.smtp_encryption}
              ico="lock-closed"
            />
          </div>
        </div>
        <div class="mt-5 flex">
          <div class="pr-3 w-64">
            <FormInput
              id="smtp_username"
              type="text"
              title="Username"
              bind:value={smtp.username}
              error={formErrors.smtp_username}
              ico="user"
            />
          </div>
          <div class="w-64">
            <FormInput
              id="smtp_password"
              type="password"
              title="Password"
              bind:value={smtp.password}
              error={formErrors.smtp_password}
              ico="finger-print"
            />
          </div>
        </div>
        <div class="flex pt-8">
          <FormButton type="submit" name="Save" color="green" />
          <div class="ml-5 mt-3">
            <span 
              on:click={() => sendTestLetter('smtp')} 
              class="cursor-pointer text-red-700"
              role="button"
              tabindex="0"
              on:keydown={(e) => { if (e.key === 'Enter' || e.key === ' ') { e.preventDefault(); sendTestLetter('smtp'); } }}
            >
              Test smtp
            </span>
          </div>
        </div>
      </form>
    </div>
  </div>
</svelte:component>

{#if drawerOpen}
  <Drawer isOpen={drawerOpen} on:close={closeDrawer} maxWidth="725px">
    {#if drawerMode === 'mail_letter_payment'}
      <Letter key="mail_letter_payment" name="mail_letter_payment" legend={letterLegend.mail_letter_payment} on:close={closeDrawer} on:send={(e) => sendTestLetter(e.detail)} />
    {:else if drawerMode === 'mail_letter_purchase'}
      <Letter key="mail_letter_purchase" name="mail_letter_purchase" legend={letterLegend.mail_letter_purchase} on:close={closeDrawer} on:send={(e) => sendTestLetter(e.detail)} />
    {/if}
  </Drawer>
{/if}
