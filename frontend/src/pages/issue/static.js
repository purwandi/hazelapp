import React from 'react';
import StoryIcon from '@atlaskit/icon-object/glyph/story/16';
import EpicIcon from '@atlaskit/icon-object/glyph/epic/16';
import TaskIcon from '@atlaskit/icon-object/glyph/task/16';
import BugIcon from '@atlaskit/icon-object/glyph/bug/16';

export const IssueTypes = [
  { icon: <StoryIcon key="story" />, label: 'Story', value: 'story' },
  { icon: <EpicIcon key="epic" />, label: 'Epic', value: 'epic' },
  { icon: <BugIcon key="bug" />, label: 'Bug', value: 'bug' },
  { icon: <TaskIcon key="task" />, label: 'Task', value: 'task' },
];

export const IssueType = ({ value }) => {
  const [elem] = IssueTypes.filter((type) => type.value === value);
  return elem.icon;
};
