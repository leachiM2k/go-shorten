import {Menu} from 'antd';
import React from 'react';
import {Link, useLocation} from 'react-router-dom';
import menuStructure from '../data/menu-structure';
import {isAllowed} from './RestrictToGroup';

const { SubMenu } = Menu;

const generateMenuItem = menuItem => {
    if (!isAllowed(menuItem.restrictToGroups)) {
        return null;
    }

    if (menuItem.children) {
        return (
            <SubMenu key={menuItem.name} title={menuItem.text}>
                {menuItem.children.map(menuItem => generateMenuItem(menuItem))}
            </SubMenu>
        );
    }

    return (
        <Menu.Item key={menuItem.name}>
            <Link to={menuItem.href}>{menuItem.text}</Link>
        </Menu.Item>
    );
}
const flattenedStructure = menuStructure.reduce((acc, x) => acc.concat({href:x.href, name:x.name}), []);

export default function CustomMenu() {
    let location = useLocation();

    let defaultSelectedKey;
    flattenedStructure.some(item => {
        if(item.href === location.pathname) {
            defaultSelectedKey = item.name;
            return true;
        }
        return false;
    });

    return (
        <Menu
            mode="inline"
            defaultSelectedKeys={defaultSelectedKey}
            style={{ height: '100%', borderRight: 0 }}
        >
            {menuStructure.map(menuItem => generateMenuItem(menuItem))}
        </Menu>
    )

}
