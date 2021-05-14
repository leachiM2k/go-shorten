import React, {Component} from 'react';
import {Button} from 'antd';

class SocialButton extends Component {
    render() {
        const { children, triggerLogin, icon, isLoggedIn } = this.props
        if (isLoggedIn) {
            return null;
        }
        return (
            <Button onClick={triggerLogin} icon={icon}>
                {children}
            </Button>
        )
    }
}

export default SocialButton;
