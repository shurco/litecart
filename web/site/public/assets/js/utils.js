export async function importTemplate(name) {
  return (await fetch(name)).text();
}

export function costFormat(cost) {
  return (Number(cost) ? (Number(cost) / 100).toFixed(2) : "0.00");
}