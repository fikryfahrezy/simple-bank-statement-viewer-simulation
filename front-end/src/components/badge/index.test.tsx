import { expect, test } from "vitest";
import { render, screen } from "@/test-utils";
import { Badge } from "./index";

test("Make sure the badge render", async () => {
  render(<Badge variant="primary">Success</Badge>);

  expect(screen.getByText("Success")).toBeVisible();
});
