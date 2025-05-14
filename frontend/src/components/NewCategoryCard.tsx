import React from 'react';
import { Link } from 'react-router-dom';
import { Card } from './ui/card';
import { Folder } from 'lucide-react';

interface CategoryCardProps {
  title: string;
  description: string;
  href: string;
  icon?: string;
}

const CategoryCard = ({ title, description, href, icon }: CategoryCardProps) => {
  return (
    <Link to={href}>
      <Card className="p-6 h-full hover:border-primary transition-colors">
        <div className="flex items-start gap-4">
          <div className="p-2 rounded-lg bg-muted">
            <Folder className="h-6 w-6" />
          </div>
          <div className="flex-1">
            <h3 className="font-semibold mb-2">{title}</h3>
            <p className="text-sm text-muted-foreground line-clamp-2">
              {description}
            </p>
          </div>
        </div>
      </Card>
    </Link>
  );
};

export default CategoryCard;
