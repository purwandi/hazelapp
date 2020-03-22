import React from 'react'
import Header from '../../../components/header'
import Octicon, { Plus as PlusIcon, PrimitiveDot as PrimitiveDotIcon } from '@primer/octicons-react'

const Item = (props) => (
  <div className="flex justify-between items-center border-b border-gray-200 py-1 text-sm">
    <div className="flex-1 truncate">
      <span className="text-gray-800 font-lighter tracking-wide">{props.body}</span>
    </div>
    <div className="">
      <div className="flex justify-end items-center">
        <div className="pl-2 text-gray-600">{props.id}</div>
        <div className="flex justify-between">
          <div className="px-2 ml-1 inline-block border border-gray-200 text-red-700 rounded-lg">
            <Octicon icon={PrimitiveDotIcon} />
            <span className="ml-1">Bug</span>
          </div>
          <div className="px-2 ml-1 inline-block border border-gray-200 text-yellow-700 rounded-full">
            <Octicon icon={PrimitiveDotIcon} />
            <span className="ml-1">Bug</span>
          </div>
        </div>
        <div className="px-2 text-gray-600 text-xs">Apr 29</div>
        <div className="">
          <img
            src="https://avatars.slack-edge.com/2020-02-13/949588721668_1ea008d5d04743ef5131_72.png"
            className="rounded-full inline-block w-6"
            alt="avatar"
          />
        </div>
      </div>
    </div>
  </div>
);

const Backlog = props => (
  <div className="mb-10">
    <div className="flex items-center text-sm">
      <h2 className="font-semibold">{props.title}</h2>
      <div className="ml-3">
        <div className="text-gray-500"> 3 issues</div>
      </div>
    </div>
    <div className="text-gray-600 text-xs">
      Implement the new weather alert system and make over 50,000+ customers
    </div>
    <div className="mt-2 mb-1">
      <Item id="XDM-1" body="Responsive Variations for `sticky-top`" />
      <Item id="XDM-30365" body="Create mixin for anchor tag (link)." />
      <Item id="XDM-30360" body="Segoe UI clipping with overflow:hidden" />
      <Item id="XDM-30343" body="HTML style reboot issue for placeholder link " />
      <Item id="XDM-19145" body="Allow CSS class strings in JS plugins to be overridden" />
      <Item id="XDM-19436" body="Tooltip - jQuery data caching issue " />
      <Item id="XDM-20219" body="Support multiple targets for tabs " />
      <Item id="XDM-30343" body="HTML style reboot issue for placeholder link " />
      <Item id="XDM-19145" body="Allow CSS class strings in JS plugins to be overridden" />
      <Item id="XDM-19436" body="Tooltip - jQuery data caching issue " />
      <Item id="XDM-20219" body="Support multiple targets for tabs " />
    </div>
    <div className="text-sm text-gray-600 px-2 py-1 cursor-pointer">
      <Octicon icon={PlusIcon} />
      <span className="ml-1">Create issues</span>
    </div>
  </div>
)

const BacklogBoardLayout = () => {
  return (
    <div className="h-screen flex flex-col">
      <Header name="Backlog" />
      <div className="flex-1 overflow-y-auto p-5">
        <Backlog title="Sprint 1" />
        <Backlog title="Backlog" />
      </div>
    </div>
  )
}

export default BacklogBoardLayout
