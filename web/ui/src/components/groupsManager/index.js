import React from 'react';
import './index.css';
import Group from '../group';

const API_URL = "localhost:5000/getGroups"

class GroupsManager extends React.Component {
    constructor(props) {
        super(props);

        this.state = {
            groups: null,
            error: null,
        }
    }
    componentDidMount() {
        this.loadData(); // Fetch data immediately
        this.interval = setInterval(() => this.loadData(), 8000); // Fetch data every 8 seconds
    }
    
    componentWillUnmount() {
        clearInterval(this.interval); // Clear the interval when the component unmounts
    }

    loadData() {
        fetch(API_URL)
          .then(response => response.json())
          .then(data => {
            this.setState({ groups: data, error: null });
          })
          .catch(error => {
            this.setState({ error: error });
          });
    }
    render(){
        const { groups, error } = this.state;

        if (error) {
          return <div>Error: {error.message}</div>;
        }
    
        if (!groups) {
          return <div>Loading...</div>;
        }
    
        return(
        <div className='groupsView'>
            {this.state.groups.map((group,index) => {
                return <Group key={group} groupName={group} groupDescription={group}/>
            })}
        </div>
        )
    }
}
export default GroupsManager;