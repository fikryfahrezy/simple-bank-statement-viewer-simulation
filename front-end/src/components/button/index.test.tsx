import { expect, test, vi } from "vitest";
import { render, screen } from "@/test-utils";
import { Button } from "./index";

test("Make sure the button clickable", async () => {
  const onClick = vi.fn();
  const { user } = render(<Button onClick={onClick}>Click Me</Button>);

  const button = screen.getByRole("button", { name: "Click Me" });
  expect(button).toBeVisible();
  await user.click(button);

  expect(onClick).toBeCalledTimes(1);
});
