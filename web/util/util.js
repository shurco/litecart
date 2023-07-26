export async function importTemplate(name) {
  const res = await fetch(name);
  const html = await res.text();
  return html;
}

export async function costFormat(cost) {
  return (cost ? (cost / 100).toFixed(2) : "0");
}