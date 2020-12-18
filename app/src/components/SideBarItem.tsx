import React from 'react';
import {Divider, Text} from "@chakra-ui/layout";
import styled from "styled-components";
import {SideBarItemHeader} from "./SideBarItemHeader";

const SideBarItemWrapper = styled.div`
  cursor: pointer;
  padding: 0 20px 0 20px;
  margin-top: 0 !important; 
`

export interface SideBarItemProps {
  title: string
  description: string;
}

export const SideBarItem: React.FC<SideBarItemProps> = (props) => {
  const {title, description} = props;
  return (
    <SideBarItemWrapper>
      <SideBarItemHeader>{title}</SideBarItemHeader>
      <Text isTruncated color="gray.500" noOfLines={2}>{description}</Text>
      <Divider orientation="horizontal"/>
    </SideBarItemWrapper>
  );
};

