import { expect, test } from "vitest";
import { render, screen } from "@/test-utils";
import { Modal } from "./index";

test("When modal is open, make sure the content defined", () => {
  render(<Modal open={true}>Modal Content</Modal>);

  expect(screen.getByText("Modal Content")).toBeDefined();
});

test("When modal is not open, make sure the content not defined", () => {
  render(<Modal open={false}>Modal Content</Modal>);

  expect(screen.queryByText("Modal Content")).not.toBeInTheDocument();
});
