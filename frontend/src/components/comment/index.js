import React from "react";

const CommentForm = () => (
  <div className="flex">
    <div className="pr-4">
      <img
        src="https://avatars.slack-edge.com/2020-02-13/949588721668_1ea008d5d04743ef5131_72.png"
        className="rounded inline-block w-10"
        alt="avatar"
      />
    </div>
    <div className="flex-1 relative">
      <textarea
        className="block px-3 py-3 w-full text-sm border border-gray-300 rounded shadow-inner"
        placeholder="Leave a comment"
        rows="4"
      ></textarea>
      <div className="absolute bottom-0 right-0 p-2">
        <button
          type="submit"
          className="bg-indigo-500 text-sm text-medium text-white px-4 py-2 rounded"
        >
          Save comment
        </button>
      </div>
    </div>
  </div>
);

const Comment = () => {
  return (
    <div className="my-6 py-6 border-t border-gray-400">
      <CommentForm />
    </div>
  );
};

export default Comment;
