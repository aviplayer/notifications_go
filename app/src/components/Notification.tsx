import React from 'react';
import { Divider, Stack } from '@chakra-ui/layout';
import { Input, InputGroup, InputRightElement } from '@chakra-ui/input';
import { Textarea } from '@chakra-ui/textarea';
import { Button } from '@chakra-ui/button';
import { Select } from '@chakra-ui/react';
import styled from 'styled-components';
import { NotificationType } from '../state/types';

const NotificationWrapper = styled.div`
  width: 80%;
`;

const options = ['email', 'sms', 'webhook'];

export const Notification: React.FC<NotificationType> = (props) => {
  const {
    type,
    title,
    description,
    template,
    email,
    pwd,
    // eslint-disable-next-line @typescript-eslint/naming-convention
    smtp_server,
    // eslint-disable-next-line @typescript-eslint/naming-convention
    smpt_port,
  } = props;
  const [show, setShow] = React.useState(false);
  const handleClick = () => setShow(!show);
  return (
    <NotificationWrapper>
      <Stack spacing={3}>
        <Input placehplder="add title" value={title} />
        <Divider />
        <Select placeholder="select notification type" value={type}>
          {options.map((val, index) => (
            <option value={index}>{val}</option>
          ))}
        </Select>
        <Divider />
        <Textarea
          placeholder="add description"
          resize="vertical"
          value={description}
        />
        <Divider />
        <Textarea
          placeholder="add template"
          resize="vertical"
          value={template}
        />
        <Divider />
        <Input placehplder="add email" value={email} />
        <Divider />
        <InputGroup size="md">
          <Input
            placehplder="add password"
            value={smtp_server}
            type={show ? 'text' : 'password'}
          />
          <InputRightElement width="4.5rem">
            <Button h="1.75rem" size="sm" onClick={handleClick}>
              {show ? 'Hide' : 'Show'}
            </Button>
          </InputRightElement>
        </InputGroup>
        <Divider />
        <InputGroup size="md">
          <Input
            placehplder="enter server"
            value={pwd}
            type={show ? 'text' : 'password'}
          />
          <InputRightElement width="2.5rem">
            <Input
              placehplder="enter port"
              value={smpt_port}
              type={show ? 'text' : 'password'}
            />
          </InputRightElement>
        </InputGroup>
      </Stack>
    </NotificationWrapper>
  );
};
