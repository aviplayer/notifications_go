import React from "react";
import {useSelector} from "react-redux";
import {getContentType} from "../state/app/app.selectors";
import {Spinner} from "@chakra-ui/spinner";
import {ContentType} from "../state/app/app.types";
import {Loader, Container, SideBarStyles, Content} from "../components";
import {SideBar} from "./SideBar";
import {Notification} from "../components/Notification";

export const AppContainer: React.FC = React.memo(() => {
  const contentType = useSelector(getContentType);
  console.log(contentType);
  return (
    <>
      {contentType === ContentType.NotStarted ? (
          <>
            <Loader>
              <Spinner size="xl"/>
            </Loader>
          </>
        ) :
        (
          <Container>
            <SideBarStyles>
              <SideBar/>
            </SideBarStyles>
            <Content>
               <Notification id={0} type={1} title={"test"}
                             description={"test"} template={"<div>test</div>>"}
                             email={"amd@test.com"} pwd={"1234"}
                             smtp_server={"text.com"} smpt_port={90}/>
            </Content>

          </Container>
        )
      }
    </>
  );

});
