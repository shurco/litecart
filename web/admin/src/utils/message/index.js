// "connextSuccess" | "connextError" | "connextWarning" | "connextInfo";
export function showMessage(message, event = "connextSuccess") {
  const eventMessage = new CustomEvent(event, {
    detail: message
  })
  dispatchEvent(eventMessage)
}
