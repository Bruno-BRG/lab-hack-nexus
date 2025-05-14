import React from 'react';
import { Link } from 'react-router-dom';
import { Card } from './ui/card';
import { ChevronRight, Folder } from 'lucide-react';

interface CategoryCardProps {
  title: string;
  description: string;
  href: string;
  icon?: string;
}

const CategoryCard = ({ title, description, href, icon }: CategoryCardProps) => {
  return (
    <Link to={href} className="block">
      <Card className="p-6 h-full transition-all duration-300 hover:shadow-lg">
        <div className="flex justify-between items-start space-x-4">
          <div className="bg-primary/10 p-3 rounded-lg">
            <Folder className="h-6 w-6 text-primary" />
          </div>
        </div>
        
        <h3 className="mt-4 text-lg font-medium">{title}</h3>
        
        <p className="mt-2 text-muted-foreground text-sm line-clamp-3">
          {description}
        </p>
        
        <div className="mt-4 flex items-center text-primary group">
          <span className="text-sm font-medium">View Category</span>
          <ChevronRight className="ml-1 h-4 w-4 transition-transform group-hover:translate-x-1" />
        </div>
      </Card>
    </Link>
  );
};

export default CategoryCard;
