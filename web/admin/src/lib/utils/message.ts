export type MessageEvent =
  | "connextSuccess"
  | "connextError"
  | "connextWarning"
  | "connextInfo";

export function showMessage(
  message: string,
  event: MessageEvent = "connextSuccess",
): void {
  const eventMessage = new CustomEvent(event, {
    detail: message,
  });
  dispatchEvent(eventMessage);
}
