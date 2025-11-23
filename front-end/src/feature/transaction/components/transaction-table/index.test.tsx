import { expect, test, vi } from "vitest";
import { render, screen } from "@/test-utils";
import { TransactionTable } from "./index";

test("Make sure all header rendered correctly", async () => {
  render(
    <TransactionTable
      rows={[]}
      sortKey={null}
      sortOrder={null}
      onSortChange={vi.fn()}
    />,
  );

  expect(screen.getByRole("columnheader", { name: "Timestamp" })).toBeVisible();
  expect(screen.getByRole("columnheader", { name: "Name" })).toBeVisible();
  expect(screen.getByRole("columnheader", { name: "Type" })).toBeVisible();
  expect(screen.getByRole("columnheader", { name: "Amount" })).toBeVisible();
  expect(screen.getByRole("columnheader", { name: "Status" })).toBeVisible();
  expect(
    screen.getByRole("columnheader", { name: "Description" }),
  ).toBeVisible();
});

test("Make sure empty rows shows 'No Data'", async () => {
  render(
    <TransactionTable
      rows={[]}
      sortKey={null}
      sortOrder={null}
      onSortChange={vi.fn()}
    />,
  );

  expect(screen.getByText("No Data")).toBeVisible();
});

test("Make shows 'Loading' when loading", async () => {
  render(
    <TransactionTable
      rows={[]}
      sortKey={null}
      sortOrder={null}
      onSortChange={vi.fn()}
      loading={true}
    />,
  );

  expect(screen.getByText("Loading")).toBeVisible();
});
