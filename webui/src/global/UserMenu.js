import {Menu} from 'antd';
import {UserOutlined} from '@ant-design/icons';
import React from 'react';
import {Link} from 'react-router-dom';

const UserMenu = () => {
    const userName = 'Unknown User';

    return (
        <Menu theme="dark" mode="horizontal" defaultSelectedKeys={['1']}>
            <Menu.Item icon={<UserOutlined/>}><Link to={"/me"}>{userName}</Link></Menu.Item>
            {/*<Menu.SubMenu icon={<UserOutlined/>} title={userName}>*/}
            {/*    <Menu.Item key="1"><Link to={"/me"}>My profile</Link></Menu.Item>*/}
            {/*    <Menu.Item key="2">Logout</Menu.Item>*/}
            {/*</Menu.SubMenu>*/}
        </Menu>
    );
};

export default UserMenu;
