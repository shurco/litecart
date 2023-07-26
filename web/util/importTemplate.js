export default async function importTemplate(name) {
  const res = await fetch(name);
  const html = await res.text();
  return html;
}