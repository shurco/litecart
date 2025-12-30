<script lang="ts">
  import '../assets/app.css';
  import { Toaster } from 'svelte-french-toast';
  import Alert from '$lib/components/Alert.svelte';
  import SvgSprite from '$lib/components/SvgSprite.svelte';
</script>

<SvgSprite />
<Toaster position="bottom-right" />
<Alert />
<slot />