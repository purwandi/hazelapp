import React from 'react';
import { Link, useRouteMatch } from 'react-router-dom';
import Octicon, {
  Plus as PlusIcon,
  PrimitiveDot as PrimitiveDotIcon,
} from '@primer/octicons-react';

const Separator = (props) => (
  <div className="flex border-b border-gray-400 px-3 py-1 bg-gray-100 text-sm">
    {/* <div className="px-2">
      <input type="checkbox" />
    </div> */}
    <div className="px-2 flex-1 tracking-wide text-gray-900">{props.title}</div>
    <div className="px-2">
      <Octicon icon={PlusIcon} />
    </div>
  </div>
);

const Item = (props) => (
  <div className="flex justify-between items-center border-b border-gray-200 px-3 py-1 text-sm">
    {/* <div className="px-2">
      <input type="checkbox" />
    </div> */}
    <div className="pl-2 w-24 text-gray-600">{props.id}</div>
    <div className="px-2 flex-1 truncate">
      <span className="text-gray-900 font-lighter tracking-wide">{props.body}</span>
    </div>
    <div className="px-1">
      <div className="flex justify-end items-center">
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

const AppIssuesLayout = (props) => {
  const { url } = useRouteMatch();

  return (
    <div className="h-screen flex flex-col">
      <header className="flex justify-between items-center py-3 border-b border-gray-200">
        <h1 className="tracking-wider font-medium px-5 flex-1">Active Issues</h1>
        <div className="px-5">
          <Link to={`${url}/create`} className="bg-indigo-500 py-1 px-2 text-white rounded text-sm">
            <Octicon icon={PlusIcon} />
          </Link>
        </div>
      </header>
      <div className="overflow-y-auto flex-1 pb-5">
        <Separator title="Ready for testing" />
        <Item id="XDM-1" body="Responsive Variations for `sticky-top`" />
        <Item id="XDM-30365" body="Create mixin for anchor tag (link)." />
        <Item id="XDM-30360" body="Segoe UI clipping with overflow:hidden" />
        <Item
          id="XDM-30344"
          body="Split variable $table-cell-padding into separate variables css"
        />
        <Item id="XDM-30343" body="HTML style reboot issue for placeholder link " />
        <Item id="XDM-19145" body="Allow CSS class strings in JS plugins to be overridden" />
        <Item id="XDM-19436" body="Tooltip - jQuery data caching issue " />
        <Item id="XDM-20219" body="Support multiple targets for tabs " />

        <Separator title="In review" />
        <Item id="XDM-30385" body="Responsive Variations for `sticky-top`" />
        <Item id="XDM-30365" body="Create mixin for anchor tag (link)." />
        <Item id="XDM-30360" body="Segoe UI clipping with overflow:hidden" />
        <Item
          id="XDM-30344"
          body="Split variable $table-cell-padding into separate variables css"
        />
        <Item id="XDM-30343" body="HTML style reboot issue for placeholder link " />
        <Item id="XDM-19145" body="Allow CSS class strings in JS plugins to be overridden" />
        <Item id="XDM-19436" body="Tooltip - jQuery data caching issue " />
        <Item id="XDM-20219" body="Support multiple targets for tabs " />

        <Separator title="In design" />
        <Item id="XDM-30385" body="Responsive Variations for `sticky-top`" />
        <Item id="XDM-30365" body="Create mixin for anchor tag (link)." />
        <Item id="XDM-30360" body="Segoe UI clipping with overflow:hidden" />
        <Item
          id="XDM-30344"
          body="Split variable $table-cell-padding into separate variables css"
        />
        <Item id="XDM-30343" body="HTML style reboot issue for placeholder link " />
        <Item id="XDM-19145" body="Allow CSS class strings in JS plugins to be overridden" />
        <Item id="XDM-19436" body="Tooltip - jQuery data caching issue " />
        <Item id="XDM-20219" body="Support multiple targets for tabs " />

        <Separator title="Todo" />
        <Item id="XDM-30385" body="Responsive Variations for `sticky-top`" />
        <Item id="XDM-30365" body="Create mixin for anchor tag (link)." />
        <Item id="XDM-30360" body="Segoe UI clipping with overflow:hidden" />
        <Item
          id="XDM-30344"
          body="Split variable $table-cell-padding into separate variables css"
        />
        <Item id="XDM-30343" body="HTML style reboot issue for placeholder link " />
        <Item id="XDM-19145" body="Allow CSS class strings in JS plugins to be overridden" />
        <Item id="XDM-19436" body="Tooltip - jQuery data caching issue " />
        <Item id="XDM-20219" body="Support multiple targets for tabs " />
      </div>
    </div>
  );
};

export default AppIssuesLayout;
