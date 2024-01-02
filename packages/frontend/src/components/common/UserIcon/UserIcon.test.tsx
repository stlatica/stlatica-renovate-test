import { render } from "@testing-library/react";
import React from "react";
import { describe, expect, test } from "vitest";

import { Story1, Story2, Story3 } from "./UserIcon.stories";

describe("snapshot", () => {
  test("snapshot1", () => {
    const { asFragment } = render(<Story1 />);
    expect(asFragment()).toMatchSnapshot();
  });
  test("snapshot2", () => {
    const { asFragment } = render(<Story2 />);
    expect(asFragment()).toMatchSnapshot();
  });
  test("snapshot3", () => {
    const { asFragment } = render(<Story3 />);
    expect(asFragment()).toMatchSnapshot();
  });
});
