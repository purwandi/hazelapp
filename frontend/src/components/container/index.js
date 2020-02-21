import styled from "styled-components";

const AppContainer = styled.div`
  height: calc(100vh - 57px);
  overflow: scroll;

  &::-webkit-scrollbar {
    width: 5px;
  }

  &::::-webkit-scrollbar-track {
    -webkit-box-shadow: inset 0 0 6px rgba(0, 0, 0, 0.3);
  }

  &::-webkit-scrollbar-thumb {
    -webkit-box-shadow: inset 0 0 6px rgba(0, 0, 0, 0.5);
  }
`;

export default AppContainer;
