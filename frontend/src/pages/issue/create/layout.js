import React, { useRef } from 'react';
import { useSetState, useClickAway } from 'react-use';
import Octicon, { PlusSmall as PlusIcon, Saved as SavedIcon } from '@primer/octicons-react';

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
      <div ref={ref} className={`relative ${state.form}`}>
        <form onSubmit={props.onFormSubmit}>
          <div className="w-8 absolute py-1 px-2">
            <Octicon icon={SavedIcon} />
          </div>
          <input
            type="text"
            className="flex-1 pr-1 pl-8 py-1 text-sm w-full border border-indigo-300 outline-none"
            placeholder="What needs to be done?"
            value={props.state.title}
            onChange={e => props.setState({ title: e.target.value })}
            ref={input => input && input.focus()}
          />
        </form>
      </div>
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
