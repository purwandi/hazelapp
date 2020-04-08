import React, { useRef } from 'react';
import { useSetState, useClickAway } from 'react-use';
import Octicon, {
  PlusSmall as PlusIcon,
  ChevronDown as ChevronDownIcon,
} from '@primer/octicons-react';
import { string, func, shape } from 'prop-types';
import { IssueTypes, IssueType } from '../static';

const Dropdown = ({ value, onChange }) => {
  const ref = useRef(null);
  const [state, setState] = useSetState({
    show: 'hidden',
  });

  useClickAway(ref, () => {
    setState({ show: 'hidden' });
  });

  return (
    <div className="relative" ref={ref}>
      <button
        type="button"
        className="cursor-pointer flex items-center focus:outline-none px-2 px-1"
        onClick={() => setState({ show: 'block' })}
      >
        <IssueType value={value} />
        <Octicon icon={ChevronDownIcon} className="ml-2 text-gray-600" />
      </button>
      <div
        className={`
          absolute bg-white mt-2 border py-1 border-gray-300 rounded
          text-sm text-gray-700 shadow w-48 ${state.show}
        `}
      >
        {IssueTypes.map((type) => {
          if (type.value !== value) {
            return (
              <div
                className="flex items-center px-2 py-1 hover:bg-gray-200 cursor-pointer"
                onClick={() => onChange(type.value)}
                key={type.value}
              >
                {type.icon}
                <span className="ml-2">{type.label}</span>
              </div>
            );
          }
        })}
      </div>
    </div>
  );
};

Dropdown.propTypes = {
  value: string.isRequired,
  onChange: func.isRequired,
};

const FormIssue = ({ onFormSubmit, onChangeType, state: { title, type }, setState }) => {
  const ref = useRef(null);
  const [form, setForm] = useSetState({
    form: 'hidden',
    button: 'block',
  });

  useClickAway(ref, () => {
    setForm({ form: 'hidden', button: 'block' });
  });

  return (
    <div className="mt-2">
      <form onSubmit={onFormSubmit} ref={ref} className={`relative ${form.form}`}>
        <div className="absolute top-0 bottom-0 flex items-center ">
          <Dropdown value={type} onChange={onChangeType} />
        </div>
        <input
          type="text"
          className="flex-1 pr-1 pl-12 py-1 text-sm w-full border border-indigo-300 outline-none"
          placeholder="What needs to be done?"
          value={title}
          onChange={(e) => setState({ title: e.target.value })}
          ref={(input) => input && input.focus()}
        />
      </form>
      <button
        type="button"
        className={`
          text-sm text-gray-600 px-2 py-1 border border-transparent w-full
          text-left hover:bg-gray-100 focus:outline-none hover:text-gray-800
          rounded ${form.button}
        `}
        onClick={() => setForm({ form: 'block', button: 'hidden' })}
      >
        <Octicon icon={PlusIcon} />
        <span className="pl-2">Create Issue</span>
      </button>
    </div>
  );
};

FormIssue.propTypes = {
  onFormSubmit: func.isRequired,
  onChangeType: func.isRequired,
  setState: func.isRequired,
  state: shape({
    title: string.isRequired,
    type: string.isRequired,
  }).isRequired,
};

export default FormIssue;
