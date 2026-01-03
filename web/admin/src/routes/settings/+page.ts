import { redirect } from '@sveltejs/kit'
import { base } from '$app/paths'
import type { PageLoad } from './$types'

export const load: PageLoad = () => {
  throw redirect(302, `${base}/settings/main`)
}
