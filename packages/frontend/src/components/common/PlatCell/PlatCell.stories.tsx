import React from "react";
import type { Story, StoryDefault } from "@ladle/react";
import { PlatCell } from "./PlatCell";

export default {
  title: "components/common/PlatCell",
} satisfies StoryDefault;

export const Story1: Story = () => {
  return (
    <PlatCell
      content="3週連続でレート下がってるんだが。"
      favoriteCount={4}
      replyCount={0}
      shareCount={1}
      userName="夏ノブチ"
      userId="@Nobuchi32384"
    />
  );
};
Story1.storyName = "1行のPlat";

export const Story2: Story = () => {
  return (
    <PlatCell
      content="nobuchiさんのパナソニックグループ プログラミングコンテスト2023（AtCoder Beginner Contest 326）での成績：4526位, パフォーマンス：670相当, レーティング：789→777 (-12) "
      favoriteCount={8}
      replyCount={1}
      shareCount={3}
      userName="冬ノブチ"
      userId="@Nobuchi79323"
    />
  );
};
Story2.storyName = "複数行のPlat";
