import React, { useRef } from 'react';
import { useSetState, useClickAway } from 'react-use';
import Octicon, {
  PlusSmall as PlusIcon,
  ChevronDown as ChevronDownIcon
} from '@primer/octicons-react';
import StoryIcon from '@atlaskit/icon-object/glyph/story/16'
import EpicIcon from '@atlaskit/icon-object/glyph/epic/16'
import TaskIcon from '@atlaskit/icon-object/glyph/task/16'
import BugIcon from '@atlaskit/icon-object/glyph/bug/16'

const Dropdown = props => {
  const ref = useRef(null);
  const [state, setState] = useSetState({
    show: 'hidden',
    options: [
      { icon: <StoryIcon key="story" />, label: 'Story', value: 'story', },
      { icon: <EpicIcon key="epic" />, label: 'Epic', value: 'epic' },
      { icon: <BugIcon key="bug" />, label: 'Bug', value: 'bug' },
      { icon: <TaskIcon key="task" />, label: 'Task', value: 'task' }
    ]
  });

  useClickAway(ref, () => {
    setState({ show: 'hidden' });
  });

  return (
    <div className="relative" ref={ref} >
      <button type="button" className="cursor-pointer flex items-center focus:outline-none px-2 px-1" onClick={() => setState({ show: 'block' })}>
        {state.options.map(type => {
          if (type.value === props.value) {
            return type.icon
          }
        })}
        <Octicon icon={ChevronDownIcon} className="ml-2 text-gray-600" />
      </button>
      <div className={`absolute bg-white mt-2 border py-1 border-gray-300 rounded text-sm text-gray-700 shadow w-48 ${state.show}`}>
        {state.options.map(type => {
          if (type.value !== props.value) {
            return (
              <div
                className="flex items-center px-2 py-1 hover:bg-gray-200 cursor-pointer"
                onClick={() => props.onChange(type.value)}
                key={type.value}
              >
                {type.icon}
                <span className="ml-2">{type.label}</span>
              </div>
            )
          }
        })}
      </div>
    </div>
  )
}

const FormIssue = props => {
  const ref = useRef(null);
  const [state, setState] = useSetState({
    form: 'hidden',
    button: 'block',
  });

  useClickAway(ref, () => {
    setState({ form: 'hidden', button: 'block' });
  });

  return (
    <div className="mt-2">
      <form onSubmit={props.onFormSubmit} ref={ref} className={`relative ${state.form}`}>
        <div className="absolute top-0 bottom-0 flex items-center ">
          <Dropdown value={props.state.type} onChange={props.onChangeType} />
        </div>
        <input
          type="text"
          className="flex-1 pr-1 pl-12 py-1 text-sm w-full border border-indigo-300 outline-none"
          placeholder="What needs to be done?"
          value={props.state.title}
          onChange={e => props.setState({ title: e.target.value })}
          ref={input => input && input.focus()}
        />
      </form>
      <button
        type="button"
        className={`text-sm text-gray-600 px-2 py-1 border border-transparent w-full text-left hover:bg-gray-100 focus:outline-none hover:text-gray-800 rounded ${state.button}`}
        onClick={() => setState({ form: 'block', button: 'hidden' })}
      >
        <Octicon icon={PlusIcon} />
        <span className="pl-2">Create Issue</span>
      </button>
    </div>
  );
};

export default FormIssue;
