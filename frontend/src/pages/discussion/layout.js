import React from 'react';

const CommentItem = () => (
  <div className="flex items-start justify-between py-1 text-sm">
    <div className="w-12 h-12 mt-1">
      <img
        src="https://avatars.slack-edge.com/2020-02-13/949588721668_1ea008d5d04743ef5131_72.png"
        className="rounded inline-block"
        alt="avatar"
      />
    </div>
    <div className="pl-2 tracking-wide flex-1">
      <a href="/" className="font-semibold">
        Member
      </a>
      <span className="ml-2 text-gray-700 text-xs">11:10 AM</span>
      <p className="text-gray-800 mt-1">
        Utilities for controlling how an element handles content that is too large for the
        container. Utilities for controlling how an element handles content that is too large for
        the container.
      </p>
    </div>
  </div>
);

const AppDiscussionLayout = () => (
  <div className="flex flex-col overflow-hidden h-screen">
    <div className="px-5">
      <header className="flex justify-between items-center border-b border-gray-200 py-2">
        <h1 className="tracking-wide text-sm">Chat</h1>
        <button
          type="button"
          className="bg-indigo-500 text-white flex justify-between items-center rounded-sm px-2 py-1"
        >
          <span className="text-xs px-2">Chat</span>
        </button>
      </header>
    </div>
    <div className="px-5 py-4 overflow-y-scroll flex-1">
      <CommentItem />
      <CommentItem />
      <CommentItem />
      <CommentItem />
      <CommentItem />
      <CommentItem />
      <CommentItem />
    </div>
    <div className="px-5 py-4">
      <input
        type="text"
        placeholder="Message #chat"
        className="w-full border border-gray-400 shadow-xs text-xs px-3 py-2 rounded-sm"
      />
    </div>
  </div>
);

export default AppDiscussionLayout;
