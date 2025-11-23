import { expect, test } from "vitest";
import { render, screen } from "@/test-utils";
import { ErrorMessage } from "./index";

test("Make sure the button clickable", async () => {
  render(<ErrorMessage>Uh oh! Something happening</ErrorMessage>);

  expect(screen.getByText("Uh oh! Something happening")).toBeVisible();
});
