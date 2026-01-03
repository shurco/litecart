import { DEFAULT_PAGE_SIZE, DEFAULT_PAGE } from '$lib/constants/pagination'

export function usePagination(initialLimit = DEFAULT_PAGE_SIZE) {
  const currentPage = $state(DEFAULT_PAGE)
  const limit = $state(initialLimit)
  const total = $state(0)

  const totalPages = $derived(Math.ceil(total / limit))

  function setPage(page: number) {
    currentPage = page
  }

  function setTotal(newTotal: number) {
    total = newTotal
  }

  function setLimit(newLimit: number) {
    limit = newLimit
  }

  return {
    get currentPage() {
      return currentPage
    },
    get limit() {
      return limit
    },
    get total() {
      return total
    },
    get totalPages() {
      return totalPages
    },
    setPage,
    setTotal,
    setLimit
  }
}
