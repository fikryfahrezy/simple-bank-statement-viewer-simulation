import { expect, test } from "vitest";
import { render, screen } from "@/test-utils";
import { Table, TableHead, TableHeader, TableRow } from "./index";

test("Check table and table head accesible", async () => {
  render(
    <Table>
      <TableHeader>
        <TableRow>
          <TableHead>Test</TableHead>
        </TableRow>
      </TableHeader>
    </Table>,
  );

  expect(screen.getByRole("table")).toBeDefined();
  expect(screen.getByRole("columnheader")).toBeDefined();
});
