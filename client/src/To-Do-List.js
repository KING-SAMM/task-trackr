import React, { Component } from 'react';
import axios from 'axios';
import { Card, Header, Form, Input, Icon } from 'semantic-ui-react';

const endpoint = 'http://localhost:9000';

class TodoList extends Component {
    constructor(props){
        super(props)

        this.state = {
            task: "",
            items: []
        };
    }

    ComponentDidMount() {
        this.getTask();
    }

    render() {
        return (
            <div>
                <div className='row'>
                    <Heaeder className="header" color="yellow">
                        To Do List
                    </Heaeder>
                </div>
                <div className="row">
                    <form onSubmit={this.onSubmit}>
                        <Input 
                            type="text"
                            name="task"
                            onChange={this.onChange}
                            value={this.state.task}
                            fluid
                            placeholder="Create task"
                            />
                    </form>
                </div>
            </div>
        )
    }
}

export default TodoList;